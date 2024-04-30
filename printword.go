package prac

import (
	"errors"
	"fmt"
)

func PrintWord(word string, asciiMap [][]string) error {
	if asciiMap == nil || len(asciiMap) != 95 || len(asciiMap[0]) != 9 {
		return errors.New("invalid asciiMap")
	}

	for i := 1; i <= 8; i++ {
		for _, char := range word {
			ind := int(char - 32)
			if ind < 0 || ind >= len(asciiMap) {
				return fmt.Errorf("character '%c' is out of range", char)
			}
			if i >= len(asciiMap[ind]) {
				return fmt.Errorf("asciiMap does not contain enough rows for character '%c'", char)
			}
			fmt.Print(asciiMap[ind][i])
		}
		fmt.Println()
	}
	return nil
}
