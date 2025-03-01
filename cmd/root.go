package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if len(os.Args) < 2 {
		fmt.Println("Error: CSV file path is required")
		fmt.Println("Usage: csv2txt [csvfile]")
		os.Exit(1)
	}

	csvFilePath := os.Args[1]

	fmt.Printf("Processing CSV file: %s\n", csvFilePath)
}
