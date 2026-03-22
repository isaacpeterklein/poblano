package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"poblano/internal/generator"
	"poblano/internal/parser"
)

func Build() {
	// Find .pob file in current directory
	matches, err := filepath.Glob("*.pob")
	if err != nil || len(matches) == 0 {
		fmt.Println("No .pob file found in current directory.")
		os.Exit(1)
	}

	pobFile := matches[0]
	if len(matches) > 1 {
		fmt.Printf("Multiple .pob files found, using: %s\n", pobFile)
	}

	content, err := os.ReadFile(pobFile)
	if err != nil {
		fmt.Printf("Error reading %s: %v\n", pobFile, err)
		os.Exit(1)
	}

	site, err := parser.Parse(string(content))
	if err != nil {
		fmt.Printf("Parse error: %v\n", err)
		os.Exit(1)
	}

	if err := generator.Build(site, "dist"); err != nil {
		fmt.Printf("Build error: %v\n", err)
		os.Exit(1)
	}
}
