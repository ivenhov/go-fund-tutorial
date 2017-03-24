package main

import (
	"bufio"
	"fmt"
	"os"
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
	in := bufio.NewReader(os.Stdin)
	num, err := 0, 0
	for err == 0 {
		num, err = readNum(in)
		fmt.Printf("Number read: %d\n", num)
	}
}
