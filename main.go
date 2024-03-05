package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n\n", err)
		os.Exit(1)
	}
	fmt.Printf("Terminal width: %d columns\n", width)
	listFilesAndFolders(".")
}

func listFilesAndFolders(dirPath string) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		return
	}

	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() {
			// Attempt to print directory names in blue
			fmt.Printf("\033[34m%s\033[0m\n", name)
		} else {
			// Print file names in default color
			fmt.Println(name)
		}
	}
}
