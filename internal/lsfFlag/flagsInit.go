package lsfFlag

import (
	"flag"
	"fmt"
)

type LsfCmds struct {
	Tg_listSize bool
	Directory   string
	Renamer     string
}

func Construct() LsfCmds {
	return LsfCmds{
		Tg_listSize: false,
		Directory:   ".",
		Renamer:     "",
	}
}

func InitFlags(args []string, lsfState *LsfCmds) {
	flagSet := flag.NewFlagSet(".", flag.ContinueOnError)

	listSize := flagSet.Bool("l", false, "List files and folders with their lengths")
	seqRename := flagSet.String("seq-rename", "", "Specify the sequence rename pattern")
	flagSet.StringVar(seqRename, "s", "", "Specify the seuqnece rename pattern (shorthand)")

	err := flagSet.Parse(args)

	fmt.Println("Sequential Rename Pattern: ", *seqRename)
	if err != nil {
		fmt.Println("Error parsing flags", err)
		return
	}

	lsfState.Tg_listSize = *listSize

	remainingArgs := flagSet.Args()
	if len(remainingArgs) > 0 {
		lsfState.Directory = remainingArgs[0]
	}
}
