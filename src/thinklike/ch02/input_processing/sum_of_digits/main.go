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

func double(dig int) int {
	dig *= 2
	if dig >= 10 {
		return 1 + dig%10
	}
	return dig
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	sum := 0
	for {
		r, _, _ := in.ReadRune()
		if r == '\n' {
			break
		}
		dig := digit(r)
		dig = double(dig)
		fmt.Printf("Doubled digit: %d", dig)
		sum += dig
		fmt.Print("\n")
	}
	fmt.Printf("Sum of all doubled digits: %d\n", sum)
}
