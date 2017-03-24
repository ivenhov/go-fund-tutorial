package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	//	var z = f
	fmt.Printf("Vals %d,%d,%d. Type %T", x, y, z, z)
}
