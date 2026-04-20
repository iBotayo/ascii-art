# ascii-art-output

A command-line tool written in Go that converts text into ASCII art and writes the result into a file using a flag.

---

# Table of Contents

1. [What This Program Does](#what-this-program-do)
2. [How to Run It](#how-to-run-it)
3. [Project Structure](#project-structure)
4. [How the Program Works](#how-the-program-works)
5. [Banner Files](#banner-files)
6. [Example Runs](#example-runs)
7. [Running Tests](#running-tests)

---

## What This Program Does

This program converts normal text into large **ASCII art characters** using one of three font banners:

- `standard`
- `shadow`
- `thinkertoy`

Unlike the basic ascii-art project, this version **writes the output to a file** instead of printing it to the terminal.

The output file is specified using the flag:

`--output=<filename.txt>`


**Example:**
`go run . --output=banner.txt "hello" standard`


The ASCII art will be written to `banner.txt`.

---

# How to Run It

Basic usage:

`go run . --output=<fileName.txt> "text" banner`


**Example:**


`go run . --output=hello.txt "Hello" standard`


Then view the output:


`cat hello.txt`


You should see something like:

```console
| | | | | |
| |__| | | |
| __ |/ _ \ |
| | | | __/ |
|| |_|_|_|
```
---

## Project Structure

```console
ascii-art-output/
│
├── main.go
│
├── ascii/
│ └── render.go
│
├── standard.txt
├── shadow.txt
├── thinkertoy.txt
│
└── ascii_test.go
```

## Explanation:

| File | Purpose |
|-----|------|
| main.go | Handles command-line arguments and writes output file |
| render.go | Reads banner files and builds ASCII art |
| banner files | Store ASCII art fonts |
| ascii_test.go | Unit tests |

---

## How the Program Works

The program follows these steps:

### 1. Read CLI arguments

Example command:


`go run . --output=banner.txt "Hello" standard`


    Arguments:

    os.Args[0] → program name
    os.Args[1] → --output=banner.txt
    os.Args[2] → Hello
    os.Args[3] → standard

---

### 2. Validate the flag

The program checks that the first argument starts with:


`--output=`


If not, it prints the usage message.

---

### 3. Extract the output filename

Example:


`--output=banner.txt`


becomes


    banner.txt


using:


`strings.TrimPrefix()`


---

### 4. Load the banner file

Example banner argument:


`standard`


becomes:


    standard.txt


The program reads the file and splits it into lines.

---

### 5. Build the ASCII character map

The banner file stores characters in blocks of **9 lines**:


    8 lines of ASCII art
    1 blank separator line


The program converts these into a map:


    character → 8 lines of ASCII art


Example:


`'A' → []string{line1,line2,line3,...line8}`


---

### 6. Render the ASCII art

The `PrintAscii` function:

1. Loops through each character of the input text
2. Looks up the ASCII art in the map
3. Builds the final output line by line
4. Returns a large string containing the full ASCII art

---

### 7. Write the output file

Instead of printing to the terminal, the program writes the result using:


    os.WriteFile()


Example:


`os.WriteFile(outputFile, []byte(result), 0644)`


---

## Banner Files

The banner files contain ASCII art representations for **all printable characters** from:

    ASCII 32 (space)
    to
    ASCII 126 (~)


Each character is stored as **8 rows of ASCII art**.

Example structure:


    (blank line)
    A row 1
    A row 2
    A row 3
    A row 4
    A row 5
    A row 6
    A row 7
    A row 8
    (blank separator)


The program reads these blocks and builds the character map.

---

## Example Runs

Generate ASCII output file:


`go run . --output=banner.txt "Hello" standard`


Check output:


`cat banner.txt`


Using another banner:


`go run . --output=banner.txt "Hello There!" shadow`


---

# Running Tests

Run all tests with:


`go test ./...`


or:


`go test`


Verbose output:


`go test -v`


Example output:


    === RUN TestReadBanner
    --- PASS: TestReadBanner
    === RUN TestBuildAsciiMap
    --- PASS: TestBuildAsciiMap
    PASS


---

## Requirements

- Go installed
- Only **standard Go packages**
- Banner files present in the project folder

