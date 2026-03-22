package cmd

import (
	"fmt"
	"os"
)

func Build(file string) {
	if err := buildSite(file); err != nil {
		fmt.Printf("Build error: %v\n", err)
		os.Exit(1)
	}
}
