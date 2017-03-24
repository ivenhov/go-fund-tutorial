package main

import (
	"fmt"
)

func main() {
	size := 5
	for i := 1; i <= size; i++ {
		count := size + 1 - i
		for i := 1; i <= count; i++ {
			fmt.Print("#")
		}
		fmt.Print("\n")
	}
}
