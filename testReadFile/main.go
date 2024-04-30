package main

import (
	"fmt"
	"path/filepath"
	"prac"
)

func main() {
	filename := "standard.txt"
	abspath, _ := filepath.Abs(filename)
	asciiMap := prac.ReadAsciiFromFile(abspath)
	if asciiMap == nil {
		fmt.Println("Error Reading The File")
		return
	}
	for _, line := range asciiMap{
		for _, char := range line {
			fmt.Println(char)
		}
		fmt.Println()
	}
}