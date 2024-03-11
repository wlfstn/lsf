package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	lFlag := flag.Bool("l", false, "List files and folders with their lengths")
	flag.Parse() // Parse the flags

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n\n", err)
		os.Exit(1)
	}
	fmt.Printf("Terminal width: %d columns\n\n", width)

	if *lFlag {
		listFilesAndFolders(".", true)
	} else {
		listFilesAndFolders(".", false)
	}
}

func listFilesAndFolders(dirPath string, showLength bool) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		return
	}

	for _, entry := range entries {
		name := entry.Name()
		nameData := ""
		if showLength {
			nameLength := len(name)
			nameData = fmt.Sprintf(" | %d", nameLength)
		}
		if entry.IsDir() {
			fmt.Printf("\033[34m%s\033[0m%s\n", name, nameData)
		} else {
			fmt.Printf("%s%s\n", name, nameData)
		}
	}
}
