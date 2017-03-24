package main

import (
	"fmt"
)

func DivMod(val, div int) (x, y int) {
	x = val / div
	y = val % div
	return
}

func main() {
	value, div := 16, 5
	res, mod := DivMod(value, div)
	fmt.Printf("Value %d divided by: %d gives %d and modulo %d", value, div, res, mod)
}
