package main

import (
	clicmd "github.com/wlfstn/lsf/internal/cliCmd"
	clidraw "github.com/wlfstn/lsf/internal/cliDraw"
)

func main() {
	listToggle := clicmd.FlagsInit()
	width := clidraw.GetCliWidth()
	clidraw.ListFilesAndFolders(".", width, listToggle)
}
