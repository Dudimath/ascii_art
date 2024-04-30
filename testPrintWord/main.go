package main

import (
	"prac"
)

func main() {
	filename := "standard.txt"
	asciiMap := prac.ReadAsciiFromFile(filename)
	str := "Martin"
	prac.PrintWord(str,asciiMap)
}
