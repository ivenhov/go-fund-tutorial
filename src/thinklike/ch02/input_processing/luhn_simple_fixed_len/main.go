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
		fmt.Printf("Got digit: %d\n", dig)
		sum += dig
	}
	fmt.Printf("Sum is: %d\n", sum)
	if sum%10 == 0 {
		fmt.Println("Checksum valid")
	} else {
		fmt.Println("Checksum INVALID")
	}
}
