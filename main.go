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
	Name string
	URLs []string
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

func findColumnIdx(headers []string, name string) (int, error) {
	name = strings.ToLower(strings.TrimSpace(name))

	for i, header := range headers {
		if strings.ToLower(strings.TrimSpace(header)) == name {
			return i, nil
		}
	}

	return -1, fmt.Errorf("column '%s' not found in CSV", name)
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

func parse(records [][]string) ([]FileData, error) {
	result := []FileData{}

	if len(records) < 1 {
		return result, nil
	}

	titleIdx, err := findColumnIdx(records[0], "title")
	if err != nil {
		return nil, err
	}

	mrIdx, err := findColumnIdx(records[0], "mr")
	if err != nil {
		return nil, err
	}

	for i := 1; i < len(records); i++ {
		record := records[i]

		if len(record) <= titleIdx || len(record) <= mrIdx {
			continue
		}

		title := record[titleIdx]
		mr := record[mrIdx]

		if title == "" || mr == "" || !containsURL(mr) {
			continue
		}

		fileName := convertToFilename(title)
		urls := strings.Split(mr, " ")

		fileData := FileData{
			Name: fileName,
			URLs: urls,
		}
		result = append(result, fileData)
	}

	return result, nil
}

func save(path string, files []FileData) error {
	for _, file := range files {
		filePath := filepath.Join(path, file.Name)
		f, err := os.Create(filePath)
		if err != nil {
			return fmt.Errorf("failed to create file %s: %w", file.Name, err)
		}

		content := strings.Join(file.URLs, "\n")
		if _, err := f.WriteString(content); err != nil {
			f.Close()
			return fmt.Errorf("failed to write to file %s: %w", file.Name, err)
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

	fileData, err := parse(records)
	if err != nil {
		fmt.Printf("Error parsing CSV: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Found %d records with mrs\n", len(fileData))

	folderPath := filepath.Dir(csvFilePath)
	if err := save(folderPath, fileData); err != nil {
		fmt.Printf("Error saving files: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Successfully saved files")
}
