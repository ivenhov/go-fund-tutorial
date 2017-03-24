package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter number: ")
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, _ := in.ReadRune()
		if r == '\n' {
			break
		}
		dig := r - '0'
		mult := dig * 2
		fmt.Printf("Got rune %#U. Decoded digit %d doubled is %d\n", r, dig, mult)
	}
}
