package main

/*
	Read vharacter by character.
	Comma-delimited positive integers
	Three decoding modes: big letters, small leters, punctuation

	18,12312,171,763,98423,1208,216,11,500,18,241,0,32,20620,27,10
*/

import (
	"bufio"
	"fmt"
	"os"
)

const (
	M_UPPER = iota
	M_LOWER
	M_PUNCT
)

func digit(r rune) int {
	dig := r - '0'
	return int(dig)
}

func readNum(in *bufio.Reader) (int, int) {
	num := 0
	for {
		r, _, _ := in.ReadRune()
		if r == ',' {
			return num, 0
		} else if r == '\n' {
			return num, -1
		} else {
			num *= 10
			num += digit(r)
		}
	}
}

func main() {
	fmt.Print("Enter comma-separated numbers: ")
	mode := M_UPPER
	in := bufio.NewReader(os.Stdin)
	num, err := 0, 0
	var char rune
	for err == 0 {
		num, err = readNum(in)
		switch mode {
		case M_UPPER:
			c := num % 27
			if c == 0 {
				mode = M_LOWER
			} else {
				char = rune('A' + c - 1)
			}
		case M_LOWER:
			c := num % 27
			if c == 0 {
				mode = M_PUNCT
			} else {
				char = rune('a' + c - 1)
			}
		case M_PUNCT:
			c := num % 9
			switch c {
			case 0:
				mode = M_UPPER
			case 1:
				char = '!'
			case 2:
				char = '?'
			case 3:
				char = ','
			case 4:
				char = '.'
			case 5:
				char = ' '
			case 6:
				char = ';'
			case 7:
				char = '"'
			case 8:
				char = '\''
			}
		}
		if char != 0 {
			fmt.Printf("%c", char)
		}
		char = 0
	}
}
