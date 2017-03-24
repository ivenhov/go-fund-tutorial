package main

import (
	"fmt"
	"math"
)

func powLimit(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	}
	return lim
}

func main() {
	val1 := powLimit(2, 4, 20)
	val2 := powLimit(2, 5, 20)
	fmt.Printf("value1: %f, value2: %f\n", val1, val2)
}
