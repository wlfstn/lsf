package clidraw

import (
	"fmt"
	"os"
	"strings"
)

const columnWidth = 19

func ListFilesAndFolders(dirPath string, Width int, showLength bool) {
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
