package lsfFlag

import (
	"flag"
	"fmt"
)

type LsfCmds struct {
	Tg_listSize bool
	Directory   string
}

func Construct() LsfCmds {
	return LsfCmds{
		Tg_listSize: false,
		Directory:   ".",
	}
}

func InitFlags(args []string, lsfState *LsfCmds) {
	flagSet := flag.NewFlagSet(".", flag.ContinueOnError)

	listSize := flagSet.Bool("l", false, "List files and folders with their lengths")

	err := flagSet.Parse(args)
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
