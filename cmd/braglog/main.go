package main

import (
	"fmt"
	"os"
	"time"

	"github.com/alex-evans/braglog/internal/editor"
	"github.com/alex-evans/braglog/internal/fileio"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("***ERROR***")
		fmt.Println("Usage: braglog <command>")
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "edit":
		err := editAndSave()
		if err != nil {
			handleError("Error editing and saving:", err)
		}
	default:
		fmt.Println(("***ERROR***"))
		fmt.Println("Unrecognized command:", command)
		os.Exit(1)
	}
}

func editAndSave() error {
	tmpfile, err := editor.LaunchEditor()
	if err != nil {
		return err
	}

	content, err := fileio.ReadFile(tmpfile.Name())
	if err != nil {
		return err
	}

	fmt.Println("Captured Brag Content:")
	fmt.Println(string(content))

	today := time.Now().Format("2006-01-02")
	filename := fmt.Sprintf("%s.md", today)

	err = fileio.SaveToFile(filename, content)
	if err != nil {
		return err
	}

	fmt.Println("Saved Brag")
	return nil
}

func handleError(msg string, err error) {
	fmt.Printf("%s %v\n", msg, err)
	os.Exit(1)
}
