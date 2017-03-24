package main

import (
	"bufio"
	"fmt"
	"os"
)

type Stats struct {
	longestNum    int
	wordsNum      int
	highestVovels int
}

func main() {
	fmt.Print("Enter a line of text: ")
	reader := bufio.NewReader(os.Stdin)

	total := new(Stats)
	var currLen, currVows int
	var r rune
	for r != '\n' {
		r, _, _ = reader.ReadRune()
		process := r == ' ' || r == '\n'
		if process {
			if currLen == 0 {
				continue
			}
			if currLen > total.longestNum {
				total.longestNum = currLen
			}
			if currVows > total.highestVovels {
				total.highestVovels = currVows
			}
			total.wordsNum++
			currLen, currVows = 0, 0
		} else {
			currLen++
			switch r {
			case 'a', 'e', 'i', 'o', 'u', 'y':
				currVows++
			}
		}
	}

	fmt.Println("Longest word: ", total.longestNum)
	fmt.Println("Highest num vowels: ", total.highestVovels)
	fmt.Println("Number of words: ", total.wordsNum)
}
