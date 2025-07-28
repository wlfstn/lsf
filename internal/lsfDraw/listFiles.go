package lsfDraw

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sys/windows"
)

type FileEntry struct {
	Name   string
	RawLen int
	IsDir  bool
}
type FileEntries []FileEntry

func (fe *FileEntries) FilesDirectory(dirPath string) {
	dirList, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading directory: %v\n", err)
		return
	}

	for _, item := range dirList {
		name := item.Name()
		fullPath := filepath.Join(dirPath, name)
		if isHiddenWindows(fullPath) {
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
