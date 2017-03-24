package main

/*
########
 ######
  ####
   ##

	row	want	offset
	1	8		0
	2	6		1
	3	4		2
	4	2		3

	want = 10 - row * 2
	offset = row - 1

symetry
	row	want	off
	1	4		0
	2	3		1
	3	2		2
	4	1		3

	off = row - 1
	want = 4 - off

*/

import (
	"fmt"
)

const (
	DUMMY = iota * 4
	SIZE  = iota * 4
)

func symbol(r, c int) {
	want := SIZE*2 + 2 - r*2
	off := r - 1
	if c > off && c <= off+want {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}

func symbol2(r, c int) {
	off := r - 1
	want := 4 - off
	if c > off && c <= off+want {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}

// using symetry
func main() {
	for r := 1; r <= 4; r++ {
		for c := 1; c <= 4; c++ {
			symbol(r, c)
		}
		for c := 4; c >= 1; c-- {
			symbol(r, c)
		}
		fmt.Print("\n")
	}
}

func __main() {
	for r := 1; r <= 4; r++ {
		for c := 1; c <= SIZE*2; c++ {
			symbol(r, c)
		}
		fmt.Print("\n")
	}
}

func _main() {
	size := SIZE * 2
	want := size
	for r := 1; want > 2; r++ {
		want = size + 2 - r*2
		off := r - 1
		for c := 1; c <= size; c++ {
			if c > off && c <= off+want {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
