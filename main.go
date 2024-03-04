package main

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func main() {
	width, height, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Printf("Error getting terminal size: %v\n", err)
		return
	}
	fmt.Println("Terminal size: %d columns, %d rows\n", width, height)
	fmt.Println("lsf")
}
