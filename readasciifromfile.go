package prac

import (
	"bufio"
	"fmt"
	"os"
)

func ReadAsciiFromFile(filename string) [][]string{
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error Opening File", err)
		return nil
	}
	var asciiLine []string
	var asciiMap [][]string
	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		line := scanner.Text()
		asciiLine = append(asciiLine, line)
		counter++
		if counter == 9 {
			asciiMap = append(asciiMap, asciiLine)
			asciiLine = nil
			counter = 0
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Errror Scanning File",err)
		return nil
	}
	return asciiMap
}