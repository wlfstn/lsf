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
	Tg_listSize    bool
	Tg_listWidth   bool
	Tg_listVersion bool
	Directory      string
	Renamer        string
	CopyDir        uint8
}

func Construct() LsfCmds {
	return LsfCmds{
		Tg_listSize:    false,
		Tg_listWidth:   false,
		Tg_listVersion: false,
		Directory:      ".",
		Renamer:        "",
		CopyDir:        0,
	}
}

func InitFlags(args []string, lsfState *LsfCmds) {
	flagSet := flag.NewFlagSet(".", flag.ContinueOnError)

	listSize := flagSet.Bool("l", false, "List files and folders with their lengths")
	listWidth := flagSet.Bool("w", false, "List the terminal width")
	listVersion := flagSet.Bool("v", false, "List version of the software")
	flagSet.BoolVar(listVersion, "version", false, "List the version of the software")

	cDir := flagSet.Bool("c", false, "Copy directory path")
	copyDir := flagSet.Bool("copy-dir", false, "Copy directory path")
	cDirW := flagSet.Bool("c:win", false, "Copy windows style directory path")
	copyDirW := flagSet.Bool("copy-dir:win", false, "Copy windows style directory path")

	seqRename := flagSet.String("seq-rename", "", "Specify the sequence rename pattern")
	flagSet.StringVar(seqRename, "s", "", "Specify the seuqnece rename pattern")

	err := flagSet.Parse(args)
	if err != nil {
		fmt.Println("Error parsing flags", err)
		return
	}

	lsfState.Tg_listSize = *listSize
	lsfState.Tg_listWidth = *listWidth
	lsfState.Tg_listVersion = *listVersion

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
