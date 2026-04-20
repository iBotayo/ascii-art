package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	// 1. Get string from command line arguments(second argument)
	// 2. Convert string to ascii art
	// 		- Open and read banner file content
	// 		-

	if len(os.Args) != 2 {
		fmt.Println("Please enter string to convert")
		fmt.Println("Usage: go run . [STRING]")
		fmt.Println("Example: go run . something")
		return
	}
	input := os.Args[1]

	words := strings.Split(input, `\n`)

	content, err := os.ReadFile("banners/standard.txt")

	if err != nil {
		log.Fatal(err)
	}

	line := strings.Split(string(content), "\n")

	for k, word := range words {
		if word == "" {
			fmt.Println()
			continue
		}

		for i := 1; i < 9; i++ {
			for j := 0; j < len(words[k]); j++ {
				start := (int(words[k][j]-32) * 9) + i
				// fmt.Print(start)
				fmt.Print(line[start])
			}
			// fmt.Println("here")
			fmt.Println()
		}
	}
}
