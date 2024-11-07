package clicmd

import "flag"

func FlagsInit() bool {
	lFlag := flag.Bool("l", false, "List files and folders with their lengths")
	flag.Parse()

	return *lFlag
}
