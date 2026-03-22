package cmd

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"poblano/internal/generator"
	"poblano/internal/parser"
)

var targetFile string

func Serve(file string) {
	targetFile = file

	if err := buildSite(file); err != nil {
		fmt.Printf("Build error: %v\n", err)
		os.Exit(1)
	}

	go watch(file)

	port := "3000"
	fmt.Printf("Serving at http://localhost:%s\n", port)
	fmt.Println("Watching for changes... (Ctrl+C to stop)")

	fs := http.FileServer(http.Dir("dist"))
	if err := http.ListenAndServe(":"+port, fs); err != nil {
		fmt.Printf("Server error: %v\n", err)
		os.Exit(1)
	}
}

func watch(file string) {
	pobFile := file
	if pobFile == "" {
		pobFile = findPobFile()
	}
	if pobFile == "" {
		return
	}

	var lastMod time.Time
	if info, err := os.Stat(pobFile); err == nil {
		lastMod = info.ModTime()
	}

	for {
		time.Sleep(500 * time.Millisecond)
		info, err := os.Stat(pobFile)
		if err != nil {
			continue
		}
		if info.ModTime().After(lastMod) {
			lastMod = info.ModTime()
			fmt.Printf("\nChange detected in %s, rebuilding...\n", filepath.Base(pobFile))
			if err := buildSite(file); err != nil {
				fmt.Printf("Build error: %v\n", err)
			} else {
				fmt.Println("Done. Refresh your browser.")
			}
		}
	}
}

func findPobFile() string {
	matches, err := filepath.Glob("*.pob")
	if err != nil || len(matches) == 0 {
		return ""
	}
	if len(matches) > 1 {
		fmt.Printf("Multiple .pob files found, using: %s\n", matches[0])
		fmt.Println("Tip: specify a file with: poblano build <file.pob>")
	}
	return matches[0]
}

func buildSite(file string) error {
	pobFile := file
	if pobFile == "" {
		pobFile = findPobFile()
	}
	if pobFile == "" {
		return fmt.Errorf("no .pob file found in current directory")
	}

	content, err := os.ReadFile(pobFile)
	if err != nil {
		return err
	}

	site, err := parser.Parse(string(content))
	if err != nil {
		return err
	}

	for _, w := range site.Warnings {
		fmt.Printf("  warning: %s\n", w)
	}

	return generator.Build(site, "dist")
}
