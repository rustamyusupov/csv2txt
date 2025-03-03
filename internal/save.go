package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func Save(path string, files []FileData) error {
	for _, file := range files {
		filePath := filepath.Join(path, file.FileName)
		f, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", file.FileName, err)
		}

		content := strings.Join(file.URLs, "\n")
		if _, err := f.WriteString(content); err != nil {
			f.Close()
			return fmt.Errorf("failed to write to file %s: %w", file.FileName, err)
		}

		f.Close()
	}

	return nil
}
