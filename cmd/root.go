package cmd

import (
	"fmt"
	"os"

	"github.com/rustamyusupov/csv2txt/internal"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("Error: CSV file path is required")
		fmt.Println("Usage: csv2txt [csvfile]")
		os.Exit(1)
	}

	csvFilePath := os.Args[1]
	fmt.Printf("Processing: %s\n", csvFilePath)

	records, err := internal.ReadCSV(csvFilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Successfully read %d lines\n", len(records))

	fileData := internal.Parse(records)
	fmt.Printf("Found %d records with mrs\n", len(fileData))
}
