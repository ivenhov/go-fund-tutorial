package main

/*
   ##
  ####
 ######
########
########
 ######
  ####
   ##

    row	want	offset
    1	2		3
    2	4		2
    3	6		1
    4	8		0
    5	8		0
    6	6		1
    7	4		2
    8 	2		3

    want=(row-1)*2+2
    off=

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

func main() {
	size := SIZE * 2
	for r := 4; r >= 1; r-- {
		for c := 1; c <= size; c++ {
			symbol(r, c)
		}
		fmt.Print("\n")
	}
	for r := 1; r <= 4; r++ {
		for c := 1; c <= size; c++ {
			symbol(r, c)
		}
		fmt.Print("\n")
	}
}
