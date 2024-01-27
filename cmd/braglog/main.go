package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/alex-evans/braglog/internal/database/sqlite"
	"github.com/alex-evans/braglog/internal/editor"
	"github.com/alex-evans/braglog/internal/fileio"
	"github.com/alex-evans/braglog/internal/hillchart"
)

func main() {
	switch runtime.GOOS {
	case "windows":
		fmt.Println("***ERROR***")
		fmt.Println("Windows is currently not supported")
		os.Exit(1)
	case "linux":
		fmt.Println("***WARNING***")
		fmt.Println("Linux not currently supported but allwowed to run")
	case "darwin":
		// Darwin / MacOS supported
	default:
		fmt.Println("***WARNING***")
		fmt.Println("Unknown Operating System - allowed to run")
	}

	if len(os.Args) < 2 {
		fmt.Println("***ERROR***")
		fmt.Println("Usage: braglog <command>")
		os.Exit(1)
	}

	err := setupDatabase()
	if err != nil {
		handleError("Error setting up database:", err)
		os.Exit(1)
	}
	defer sqlite.Close()

	command := os.Args[1]

	switch command {
	case "init":
		// testing init db
	case "edit":
		err := editAndSave()
		if err != nil {
			handleError("Error editing and saving:", err)
		}
	case "hill":
		err := hillchart.GenerateHillChart(50, "test")
		if err != nil {
			handleError("Error hill chart generating:", err)
		}
	default:
		fmt.Println(("***ERROR***"))
		fmt.Println("Unrecognized command:", command)
		os.Exit(1)
	}
}

func setupDatabase() error {
	executablePath, err := os.Executable()
	if err != nil {
		handleError("Error setting Executable Path:", err)
	}

	databaseDir := filepath.Join(filepath.Dir(executablePath), "data")
	err = os.MkdirAll(databaseDir, 0755)
	if err != nil {
		handleError("Error creating data directory:", err)
	}

	databasePath := filepath.Join(databaseDir, "braglog.db")

	err = sqlite.Init(databasePath)
	if err != nil {
		return err
	}

	err = sqlite.MigrateDatabase(sqlite.GetDB())
	if err != nil {
		sqlite.Close()
		return err
	}

	fmt.Println("Database setup and migrations completed successfully")
	return nil
}

func editAndSave() error {
	tmpfile, err := editor.LaunchEditor()
	if err != nil {
		return err
	}
	defer os.Remove(tmpfile.Name())

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
