package fileio

import (
	"fmt"
	"os"
	"path/filepath"
)

func ReadFile(filename string) ([]byte, error) {
	fmt.Printf("Reading File - %s", filename)
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': %v", filename, err)
	}
	return content, nil
}

func SaveToFile(filename string, content []byte) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf(("failed to find home direcotry: %v"), err)
	}

	appDir := filepath.Join(homeDir, ".braglog")
	err = os.MkdirAll(appDir, 0700)
	if err != nil {
		return fmt.Errorf(("failed to create directory '%s': %v"), appDir, err)
	}

	bragFilepath := filepath.Join(appDir, filename)

	err = os.WriteFile(bragFilepath, content, 0644)
	if err != nil {
		return fmt.Errorf(("failed to save file '%s': %v"), filename, err)
	}

	return nil
}
