# CSV to Text Files

A simple command-line tool that converts CSV data into separate text files, extracting URLs from the "mr" column and using the "title" column for filenames.

## Usage

```bash
csv2txt [csvfile]
```

## Features

- Processes CSV files with "title" and "mr" columns
- Creates individual text files named after each title (converted to lowercase with hyphens)
- Saves URLs as line-separated content in the text files
- Places output files in the same directory as the input CSV
