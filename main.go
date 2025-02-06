package main

import (
	"github.com/wlfstn/lsf/internal/clicmd"
	"github.com/wlfstn/lsf/internal/clidraw"
)

func main() {
	listToggle := clicmd.FlagsInit()
	width := clidraw.GetCliWidth()
	clidraw.ListFilesAndFolders(".", width, listToggle)
}
