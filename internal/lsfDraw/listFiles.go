package lsfDraw

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
)

const columnWidth = 19
const colPadding = 2

type FileEntry struct {
	Name   string
	RawLen int
	IsDir  bool
}

func isHiddenWindows(path string) bool {
	ptr, err := windows.UTF16PtrFromString(path)
	if err != nil {
		return false
	}
	attrs, err := windows.GetFileAttributes(ptr)
	if err != nil {
		return false
	}
	return attrs&windows.FILE_ATTRIBUTE_HIDDEN != 0 || attrs&windows.FILE_ATTRIBUTE_SYSTEM != 0
}

func DynamicListFiles(dirPath string, width int, showLength bool) {
	dirList, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		return
	}

	var entries []FileEntry

	for _, item := range dirList {
		name := item.Name()
		fullPath := filepath.Join(dirPath, name)
		if isHiddenWindows(fullPath) {
			continue //continue in go exits the loop iteration
		}

		rawLen := len(name)
		if showLength {
			nameData := fmt.Sprintf(" (%d)", rawLen)
			name += nameData
			rawLen += len(nameData)
		}
		entries = append(entries, FileEntry{
			Name:   name,
			RawLen: rawLen,
			IsDir:  item.IsDir(),
		})
	}

	if len(entries) == 0 {
		return
	}

	colPadding := 2
	numCols := 1
	var numRows int
	var colWidths []int

	// Determine how many columns fit
	for tryCols := 1; tryCols <= len(entries); tryCols++ {
		rows := (len(entries) + tryCols - 1) / tryCols
		widths := make([]int, tryCols)

		// Row-first layout: i = row*numCols + col
		for row := range rows {
			for col := 0; col < tryCols; col++ {
				idx := row*tryCols + col
				if idx >= len(entries) {
					continue
				}
				entry := entries[idx]
				if entry.RawLen > widths[col] {
					widths[col] = entry.RawLen
				}
			}
		}

		total := 0
		for _, w := range widths {
			total += w + colPadding
		}

		if total > width {
			break
		}

		numCols = tryCols
		numRows = rows
		colWidths = widths
	}

	for row := 0; row < numRows; row++ {
		for col := 0; col < numCols; col++ {
			idx := row*numCols + col
			if idx >= len(entries) {
				continue
			}

			entry := entries[idx]
			padding := colWidths[col] - entry.RawLen + colPadding

			if entry.IsDir {
				colored := fmt.Sprintf("\033[34m%s\033[0m", entry.Name)
				fmt.Printf("%s%*s", colored, padding, "")
			} else {
				fmt.Printf("%-*s", colWidths[col]+colPadding, entry.Name)
			}
		}
		fmt.Println()
	}
}
