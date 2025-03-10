# MRKeeper

A simple tool that extracts URLs from XLSX files and saves them into separate text files.

## What it does

- Reads an Excel (.xlsx) file
- Finds "title" and "mr" columns
- Creates text files named after each title (lowercase with hyphens)
- Places all URLs from the "mr" column into the corresponding text file
- Saves files in the same directory as the input file

## Usage

```bash
mrkeeper path/to/file.xlsx
```
