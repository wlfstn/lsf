package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term" // This package is not available in the standard library
)

const columnWidth = 19

func main() {
	lFlag := flag.Bool("l", false, "List files and folders with their lengths")
	flag.Parse()

	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n\n", err)
		os.Exit(1)
	}
	fmt.Printf("Terminal width: %d columns\n\n", width)

	if *lFlag {
		listFilesAndFolders(".", width, true)
	} else {
		listFilesAndFolders(".", width, false)
	}
}

func listFilesAndFolders(dirPath string, Width int, showLength bool) {
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		return
	}

	termWidth := Width - 20
	lengthRemaining := termWidth

	for _, entry := range entries {
		name := entry.Name()
		nameLength := len(name)

		if nameLength >= columnWidth {
			name = name[:columnWidth-3] + ".."
			nameLength = len(name)
		}

		columnSpaces := columnWidth - nameLength
		columnGap := strings.Repeat(" ", columnSpaces)

		if lengthRemaining < columnWidth {
			fmt.Println()
			lengthRemaining = termWidth
		} else {
			lengthRemaining -= columnWidth
		}

		nameData := ""
		if showLength {
			nameData = fmt.Sprintf(" | %d", nameLength)
		}

		if entry.IsDir() {
			fmt.Printf("\033[34m%s\033[0m%s%s", name, nameData, columnGap)
		} else {
			fmt.Printf("%s%s%s", name, nameData, columnGap)
		}
	}
}
