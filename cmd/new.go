package cmd

import (
	"fmt"
	"os"
)

func New(name string) {
	filename := name + ".pob"

	if _, err := os.Stat(filename); err == nil {
		fmt.Printf("File %s already exists.\n", filename)
		os.Exit(1)
	}

	template := `config
primary #4f46e5
accent #7c3aed
font Inter
dark-mode false
site-name ` + name + `
logo
favicon ` + `

header
home about projects contact

home
hero
Welcome to ` + name + `
Built with Poblano

card
Getting Started
Edit this .pob file to customize your site. Run "poblano build" to generate it.

about
card
About Me
Tell your story here.

card
Skills
List your skills here.

projects
card
Project One
Description of your first project.

card
Project Two
Description of your second project.

contact
card
Get In Touch
Your contact information goes here.

footer
` + name + ` — Built with Poblano
`

	if err := os.WriteFile(filename, []byte(template), 0644); err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created %s\n", filename)
	fmt.Printf("Edit it, then run: poblano build\n")
}
