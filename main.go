package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type FileData struct {
	FileName string
	URLs     []string
}

func read(filePath string) ([][]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV: %w", err)
	}

	return records, nil
}

func containsURL(line string) bool {
	urlRegex := regexp.MustCompile(`https?://[^\s]+`)
	return urlRegex.MatchString(line)
}

func convertToFilename(title string) string {
	result := strings.ToLower(title)
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, " ", "-")
	result = result + ".txt"
	return result
}

func parse(records [][]string) []FileData {
	result := []FileData{}

	for _, record := range records {
		title := record[5]
		mr := record[7]

		if len(record) < 10 || title == "" || mr == "" || !containsURL(mr) {
			continue
		}

		fileName := convertToFilename(title)
		urls := strings.Split(mr, " ")

		fileData := FileData{
			FileName: fileName,
			URLs:     urls,
		}
		result = append(result, fileData)
	}

	return result
}

func save(path string, files []FileData) error {
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

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Error: CSV file path is required")
		fmt.Println("Usage: csv2txt [csvfile]")
		os.Exit(1)
	}

	csvFilePath := os.Args[1]
	fmt.Printf("Processing: %s\n", csvFilePath)

	records, err := read(csvFilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully read %d lines\n", len(records))

	fileData := parse(records)
	fmt.Printf("Found %d records with mrs\n", len(fileData))

	folderPath := filepath.Dir(csvFilePath)
	if err := save(folderPath, fileData); err != nil {
		fmt.Printf("Error saving files: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully saved to txt files")
}
