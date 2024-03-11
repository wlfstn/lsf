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
	fmt.Printf("Terminal width: %d columns\n\n", width)
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
		nameLength := len(name)
		if entry.IsDir() {
			fmt.Printf("\033[34m%s\033[0m | (%d) \n", name, nameLength)
		} else {
			fmt.Printf("%s | (%d)\n", name, nameLength)
		}
	}
}
