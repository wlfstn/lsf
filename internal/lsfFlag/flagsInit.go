package lsfFlag

import (
	"flag"
	"fmt"
)

const (
	NoCopy uint8 = iota
	CopyStandard
	CopyWindows
)

type LsfCmds struct {
	Tg_listSize  bool
	Tg_listWidth bool
	Directory    string
	Renamer      string
	CopyDir      uint8
}

func Construct() LsfCmds {
	return LsfCmds{
		Tg_listSize:  false,
		Tg_listWidth: false,
		Directory:    ".",
		Renamer:      "",
		CopyDir:      0,
	}
}

func InitFlags(args []string, lsfState *LsfCmds) {
	flagSet := flag.NewFlagSet(".", flag.ContinueOnError)

	listSize := flagSet.Bool("l", false, "List files and folders with their lengths")
	listWidth := flagSet.Bool("dw", false, "List files and folders with their lengths")

	cDir := flagSet.Bool("c", false, "List files and folders with their lengths")
	copyDir := flagSet.Bool("copy-dir", false, "List files and folders with their lengths")
	cDirW := flagSet.Bool("c:win", false, "List files and folders with their lengths")
	copyDirW := flagSet.Bool("copy-dir:win", false, "List files and folders with their lengths")

	seqRename := flagSet.String("seq-rename", "", "Specify the sequence rename pattern")
	flagSet.StringVar(seqRename, "s", "", "Specify the seuqnece rename pattern (shorthand)")

	err := flagSet.Parse(args)
	if err != nil {
		fmt.Println("Error parsing flags", err)
		return
	}

	lsfState.Tg_listSize = *listSize
	lsfState.Tg_listWidth = *listWidth

	if *cDir || *copyDir {
		lsfState.CopyDir = CopyStandard
	} else if *cDirW || *copyDirW {
		lsfState.CopyDir = CopyWindows
	} else {
		lsfState.CopyDir = NoCopy
	}

	remainingArgs := flagSet.Args()
	if len(remainingArgs) > 0 {
		lsfState.Directory = remainingArgs[0]
	}
}
