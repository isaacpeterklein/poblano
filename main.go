package main

import (
	"fmt"
	"os"

	"poblano/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: poblano <command> [args]")
		fmt.Println("Commands:")
		fmt.Println("  new <name>   Create a new .pob file")
		fmt.Println("  build        Build the site from a .pob file")
		fmt.Println("  serve        Build, serve locally, and watch for changes")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("Usage: poblano new <name>")
			os.Exit(1)
		}
		cmd.New(os.Args[2])
	case "build":
		file := ""
		if len(os.Args) >= 3 {
			file = os.Args[2]
		}
		cmd.Build(file)
	case "serve":
		file := ""
		if len(os.Args) >= 3 {
			file = os.Args[2]
		}
		cmd.Serve(file)
	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}
