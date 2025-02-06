package main

import (
	"github.com/wlfstn/lsf/internal/lsfDraw"
	"github.com/wlfstn/lsf/internal/lsfFlag"
)

func main() {
	lsfState := lsfFlag.Construct()
	lsfFlag.InitFlags([]string{}, &lsfState)
	width := lsfDraw.GetCliWidth()
	lsfDraw.ListFilesAndFolders(".", width, lsfState.Tg_listSize)
}
