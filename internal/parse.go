package internal

import (
	"regexp"
	"strings"
)

type FileData struct {
	FileName string
	URLs     []string
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

func Parse(records [][]string) []FileData {
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
