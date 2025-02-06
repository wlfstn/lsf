package lsfDraw

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetCliWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting terminal size: %v\n\n", err)
		os.Exit(1)
	}
	fmt.Printf("Terminal width: %d columns\n\n", width)

	return width
}
