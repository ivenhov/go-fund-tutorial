package main

import (
	"fmt"
	"io"
	"strconv"
	"thinklike/ch02/input_processing"
)

func main() {
	dec2bin()
	bin2dec()
}

func bin2dec() {
	fmt.Print("Enter binary number: ")
	dr := input_processing.NewDigitReader()
	dr.SkipWhitespace = true
	digs := 0
	num := 0
	for {
		dig, err := dr.ReadBinaryDigit()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		digs++
		if digs > 32 {
			panic("Too many bits " + strconv.Itoa(digs) + " > 32")
		}
		num <<= 1
		num |= dig
	}
	fmt.Printf("Decimal number: %d\n", num)
}

func dec2bin() {
	fmt.Print("Enter decimal number: ")
	dr := input_processing.NewDigitReader()
	num := dr.ReadDecimalNumber()
	bits := 0
	mask := 1 << 31
	for i := 1; i <= 32; i++ {
		if num&mask == 0 {
			fmt.Print("0")
		} else {
			fmt.Print("1")
		}
		bits++
		switch bits {
		case 4:
			fmt.Print(" ")
		case 8:
			fmt.Print("  ")
			bits = 0
		}
		mask >>= 1
	}
	fmt.Print("\n")
}
