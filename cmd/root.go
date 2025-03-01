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

	// Read and process the CSV file
	records, err := internal.ReadCSV(csvFilePath)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	// Display the CSV content
	// fmt.Println("CSV contents:")
	// for i, record := range records {
	// 	fmt.Printf("Row %d: %v\n", i, record)
	// }

	fmt.Printf("Successfully read %d lines\n", len(records))
}
