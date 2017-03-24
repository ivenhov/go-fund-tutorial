package main

/*

#		     #
 ##        ##
  ###    ###
   ########
   ########
  ###    ###
 ##        ##
#            #

Quarter

row	want	off
1	1		0
2	2		1
3	3		2
4	4		3
*/

import (
	"fmt"
)

func symbol(r, c int) {
	want := r
	off := r - 1
	if c > off && c <= want+off {
		fmt.Print("#")
	} else {
		fmt.Print(" ")
	}
}

func main() {
	for r := 1; r <= 4; r++ {
		for c := 1; c <= 7; c++ {
			symbol(r, c)
		}
		for c := 7; c >= 1; c-- {
			symbol(r, c)
		}
		fmt.Print("\n")
	}
	for r := 4; r >= 1; r-- {
		for c := 1; c <= 7; c++ {
			symbol(r, c)
		}
		for c := 7; c >= 1; c-- {
			symbol(r, c)
		}
		fmt.Print("\n")
	}
}
