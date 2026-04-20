package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Ensure the minimum required arguments are provided:
	// Argument 1: --color=COLOR, Argument 2: SUBSTRING, Argument 3: TEXT
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . --color=COLOR SUBSTRING  TEXT")
		return
	}
	colorArg := os.Args[1]

	var text string
	var substring string

	if len(os.Args) == 4 {
		// With substring
		substring = os.Args[2]
		text = os.Args[3]
	} else if len(os.Args) == 3 {
		// No substring → color whole text
		text = os.Args[2]
		substring = text
	} else {
		fmt.Println("Usage:")
		fmt.Println("go run . --color=COLOR SUBSTRING TEXT")
		fmt.Println("or")
		fmt.Println("go run . --color=COLOR TEXT")
		return
	}
	// Extract the color name by removing the "--color=" prefix
	colorName := strings.TrimPrefix(colorArg, "--color=")

	// Map of supported color names to their ANSI escape codes
	colors := map[string]string{
		"red":    "\033[31m",
		"green":  "\033[32m",
		"blue":   "\033[34m",
		"yellow": "\033[33m",
	}

	// Look up the requested color; exit if it's not supported
	color, ok := colors[colorName]
	if !ok {
		fmt.Println("Unknown color:", colorName)
		return
	}

	// ANSI reset code to stop colorizing after the substring
	reset := "\033[0m"

	// Read the ASCII art font file (standard.txt)
	data, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	// Split the font file into individual lines
	lines := strings.Split(string(data), "\n")

	// Split the input text by the literal "\n" to handle multi-line input
	words := strings.Split(text, "\\n")

	printedBlank := false

	// Iterate over each word/line segment
	for w, word := range words {
		// Handle empty segments (i.e., blank lines from "\n" in input)
		if word == "" {
			// Print a blank line only if it's not the last segment and not already printed
			if w != len(words)-1 && !printedBlank {
				fmt.Println()
				printedBlank = true
			}
			continue
		}

		printedBlank = false

		// Find all occurrences of the substring within the current word
		// and store their start and end positions
		var substrLoc [][]int
		start := 0
		end := len(word)

		for {
			// Search for the substring starting from the current position
			substrStart := strings.Index(word[start:end], substring)
			if substrStart == -1 {
				break // No more occurrences found
			}

			// Calculate the actual position in the full word
			realStart := start + substrStart
			realEnd := realStart + len(substring)

			// Store the location of this occurrence
			substrLoc = append(substrLoc, []int{realStart, realEnd})

			// Move start forward to search for next occurrence
			start = realEnd
		}

		// Print each of the 8 ASCII art rows for the current word
		for i := 1; i <= 8; i++ {
			// Iterate over each character in the word
			for pos, char := range word {
				// Skip characters outside the printable ASCII range (32-126)
				if char < 32 || char > 126 {
					continue
				}

				// Calculate the line index in the font file for this character and row
				// Each character occupies 9 lines in the font file
				index := (int(char-32) * 9) + i

				// Check if this character's position falls within any substring occurrence
				colored := false
				for _, loc := range substrLoc {
					if pos >= loc[0] && pos < loc[1] {
						colored = true
						break
					}
				}

				// Print the character's ASCII row with or without color
				if colored {
					fmt.Print(color + lines[index] + reset) // Colorized
				} else {
					fmt.Print(lines[index]) // Normal
				}
			}
			fmt.Println() // Move to next row after printing all characters in a row
		}

		// Print a blank line between words/segments (except after the last one)
		if w != len(words)-1 {
			fmt.Println()
		}
	}
}
