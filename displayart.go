package prac



import (
	"fmt"
	"strings"
)

func PrintArt(word string, asciiMap [][]string) error {
	// Validate input parameters
	if word == "" {
		fmt.Print()
		return nil
	}

	if word == "\n" {
		fmt.Println()
		return nil
	}

	if HandleLn(word) && len(word) > 1 {
		word = word[:len(word)-1]
	}
	lines := strings.Split(word, "\n")
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		err := PrintWord(line, asciiMap)
		if err != nil {
			return err
		}
	}
	return nil
}

func HandleLn(word string) bool {
	for _, char := range word {
		if char != '\n' {
			return false
		}
	}
	return true
}