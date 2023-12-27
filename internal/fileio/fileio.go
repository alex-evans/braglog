package fileio

import (
	"fmt"
	"os"
)

func ReadFile(filename string) ([]byte, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file '%s': %v", filename, err)
	}
	return content, nil
}

func SaveToFile(filename string, content []byte) error {
	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		return fmt.Errorf(("failed to save file '%s': %v"), filename, err)
	}
	return nil
}
