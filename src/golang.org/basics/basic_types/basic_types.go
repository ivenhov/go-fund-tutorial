package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe     bool       = false
	MaxInt64 uint64     = 1<<64 - 1
	z        complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	const f = "%T, %v\n"

	var y complex128

	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt64, MaxInt64)
	fmt.Printf(f, y, y)
	fmt.Printf(f, z, z)
}
