package main

import (
	"fmt"
)

func main() {
	max := 5
	for i := 1; i <= max; i++ {
		fmt.Print(max + 1 - i)
		fmt.Print("\n")
	}
}
