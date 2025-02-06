package main

import (
	"os"

	"github.com/wlfstn/lsf/internal/lsfDraw"
	"github.com/wlfstn/lsf/internal/lsfFlag"
)

func main() {
	lsfState := lsfFlag.Construct()
	lsfFlag.InitFlags(os.Args[1:], &lsfState)
	width := lsfDraw.GetCliWidth()
	lsfDraw.ListFilesAndFolders(lsfState.Directory, width, lsfState.Tg_listSize)
}
