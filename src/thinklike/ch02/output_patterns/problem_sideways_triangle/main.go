package main

import (
	"fmt"
	"math"
)

func line(len int) {
	for i := 1; i <= len; i++ {
		fmt.Print("#")
	}
}

func main() {
	size := 4
	for i := 1; i <= size*2; i++ {
		abs := size - int(math.Abs(float64(i-size)))
		line(abs)
		fmt.Print("\n")
	}
}
