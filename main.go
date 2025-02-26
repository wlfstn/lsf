package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/wlfstn/lsf/internal/lsfDraw"
	"github.com/wlfstn/lsf/internal/lsfFlag"
)

func CopyToClipboard(text string) error {
	switch runtime.GOOS {
	case "windows": // Windows
		cmd := exec.Command("cmd", "/c", "echo|set /p="+text+"|clip")
		return cmd.Run()

	case "linux": // Linux
		cmd := exec.Command("xclip", "-selection", "clipboard")
		in, err := cmd.StdinPipe()
		if err != nil {
			return err
		}
		if _, err := in.Write([]byte(text)); err != nil {
			return err
		}
		in.Close()
		return cmd.Run()

	case "darwin": // macOS
		cmd := exec.Command("pbcopy")
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		in, err := cmd.StdinPipe()
		if err != nil {
			return err
		}
		if _, err := in.Write([]byte(text)); err != nil {
			return err
		}
		in.Close()
		return cmd.Run()

	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

func main() {
	lsfState := lsfFlag.Construct()
	lsfFlag.InitFlags(os.Args[1:], &lsfState)

	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if lsfState.CopyDir == lsfFlag.CopyStandard {
		normalDir := filepath.ToSlash(dir)
		fmt.Printf("Current Directory: %s", normalDir)
		CopyToClipboard(normalDir)
	} else if lsfState.CopyDir == lsfFlag.CopyWindows {
		fmt.Printf("Current Directory: %s", dir)
		CopyToClipboard(dir)
	} else {
		width := lsfDraw.GetCliWidth()

		if lsfState.Tg_listWidth {
			fmt.Printf("Terminal width: %d columns\n\n", width)
		}
		lsfDraw.ListFilesAndFolders(lsfState.Directory, width, lsfState.Tg_listSize)
	}
}
