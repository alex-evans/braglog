package editor

import (
	"fmt"
	"os"
	"os/exec"
)

func LaunchEditor() (*os.File, error) {
	tmpfile, err := os.CreateTemp("", "braglog-editor-*.md")
	if err != nil {
		return nil, fmt.Errorf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	cmd := exec.Command("vim", tmpfile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to launch Vim: %v", err)
	}

	return tmpfile, nil
}
