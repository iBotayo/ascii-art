package main

import (
	"ascii-art-output/ascii"
	"fmt"
	"os"
	"strings"
)

func main() {

	/*
		Valid usages:

		go run . --output=<output.txt> "text" banner
	*/

	// Confirm 
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println()
		fmt.Println("EX: go run . --output=<filename.txt> something standard")
		return
	}

	// output := os.Args[1]
	// input := os.Args[2]
	// bannerFile := os.Args[3] + ".txt"

	// Instead of hardcoding variables, we declare them to specifically handle them for each case
	var input string
	var bannerFile string = "standard.txt"
	var outputFile string

	// Case 1: go run . "Hello"
	if len(os.Args) == 2 {
		input = os.Args[1]
	}

	// Case 2: go run . "Hello" shadow
	if len(os.Args) ==  3 {
		input = os.Args[1]
		bannerFile = os.Args[2] + ".txt"
	}

	// Case 3: go run . --output=file.txt "Hello" standard
	if len(os.Args) == 4 {
		flag := os.Args[1]


	if !strings.HasPrefix(flag, "--output=") {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println("EX: go run . --output=<filename.txt> something standard")
		return
	}

	// extract the output file name
	outputFile = strings.TrimPrefix(flag, "--output=")
	input = os.Args[2]
	bannerFile = os.Args[3] + ".txt"

	}
	
	// read banner
	bannerLines, err := ascii.ReadBanner(bannerFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Build ASCII Map
	asciiMap := ascii.BuildAsciiMap(bannerLines)

	// Generate ASCII art
	result := ascii.PrintAscii(input, asciiMap)

	// If output flag used -> write file
	if outputFile != "" {
		err = os.WriteFile(outputFile, []byte(result), 0644)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	// Otherwise print to terminal
	fmt.Print(result)
}
