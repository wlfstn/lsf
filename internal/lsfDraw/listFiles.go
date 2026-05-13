//go:build !windows
// +build !windows

package lsfDraw

import (
	"fmt"
	"os"
	"strings"
)

func (fe *FileEntries) FilesDirectory(dirPath string) {
	dirList, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		return
	}

	for _, item := range dirList {
		name := item.Name()
		if isHiddenUnix(name) {
			continue
		}

		rawLen := len(name)
		*fe = append(*fe, FileEntry{
			Name:   name,
			RawLen: rawLen,
			IsDir:  item.IsDir(),
		})
	}
}

func isHiddenUnix(name string) bool {
	return strings.HasPrefix(name, ".")
}
