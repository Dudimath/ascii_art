package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main()  {
	content, _ := os.Open("standard.txt")
	scanner := bufio.NewScanner(content)
	ascii := make(map[byte][]string)
	counter := 0
	var char byte = ' '
	
	for scanner.Scan() {
		if counter != 8 {
			ascii[char] = append(ascii[char], scanner.Text())
			counter++
		} else {
			char++
			counter = 0
		}
	}
	args := strings.Split(os.Args[1], "\\n")
	for _, word := range args {
		if word == "" {
			fmt.Println()
		} else {
			for i:= 0 ; i <8 ; i++ {
				for _, lin := range word {
					if lin == '\n' {
						fmt.Println()
					} else {
						fmt.Printf("%s", ascii[byte(lin)][i])
					}
				}
				fmt.Println()
			}
		}
	}
}