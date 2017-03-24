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
		fmt.Printf("Entered: %#U\n", r)
		if r == '\n' {
			break
		}
	}
}
