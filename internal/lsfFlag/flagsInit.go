package lsfFlag

import "flag"

type LsfCmds struct {
	Tg_listSize bool
}

func Construct() LsfCmds {
	return LsfCmds{Tg_listSize: false}
}

func InitFlags(args []string, lsfState *LsfCmds) {
	listSize := flag.Bool("l", false, "List files and folders with their lengths")
	flag.Parse()

	lsfState.Tg_listSize = *listSize
}
