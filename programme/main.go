package main

import (
	"flag"
	"fmt"
	"os"
	"prac"
	"strings"
	"path/filepath"
)

func main() {
// Define flags
var (
	shadow     bool
	thinkerToy bool
)
flag.BoolVar(&shadow, "shadow", false, "Use shadow.txt for ASCII art")
flag.BoolVar(&thinkerToy, "thinkertoy", false, "Use thinkertoy.txt for ASCII art")

// Parse command-line flags
flag.Parse()


	// Extract input string
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Provide Input:")
		return
	}

	// Check if the third argument is "-shadow"
	if len(args) >= 2 && args[1] == "-shadow" {
		shadow = true
	} 
	if len(args) >= 2 && args[1] == "-thinkertoy"{
		thinkerToy = true
	}

	// Determine filename based on flag
	filename := "standard.txt"
	if shadow {
		filename = "shadow.txt"
	} else if thinkerToy{
		filename = "thinkertoy.txt"
	}

	// Get absolute path of the file
	abspath, err := filepath.Abs(filename)
	if err != nil {
		fmt.Println("Error getting absolute path:", err)
		return
	}

	// Read ASCII art from the file
	asciiMap := prac.ReadAsciiFromFile(abspath)
	if asciiMap == nil {
		fmt.Println("Error reading ASCII from file")
		return
	}

	// Replace "\n" with actual newline character
	str := strings.ReplaceAll(args[0], "\\n", "\n")

	// Print ASCII art
	err = prac.PrintArt(str, asciiMap)
	if err != nil {
		fmt.Println("Error printing ASCII art:", err)
	}
}