package main

import (
	"fmt"
)

func symbol(r, c int) {
	want := r
	off := 3 - want
	if c > off && c <= want+off {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}

func main() {
	rm := 3
	cm := 3
	for r := 1; r <= rm; r++ {
		for c := 1; c <= cm; c++ {
			symbol(r, c)
		}
		for c := cm; c >= 1; c-- {
			symbol(r, c)
		}
		fmt.Print("\n")
	}
	for r := rm; r >= 1; r-- {
		for c := 1; c <= cm; c++ {
			symbol(r, c)
		}
		for c := cm; c >= 1; c-- {
			symbol(r, c)
		}
		fmt.Print("\n")
	}
}
