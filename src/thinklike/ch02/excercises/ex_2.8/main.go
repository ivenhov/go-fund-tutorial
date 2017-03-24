package main

/*
	Reading binary, decimal or hexadecimal numbers and displaying them as
	binary, decimal or hexadecimal depending on user selection

*/
import (
	"fmt"
	"thinklike/ch02/input_processing"
	"thinklike/nums"
)

func main() {

	fmt.Print("Select base of the number 2, 10, 16: ")

	dr := input_processing.NewDigitReader()
	base := dr.ReadDecimalNumber()

	fmt.Printf("Enter base-%d number: ", base)

	num := dr.ReadNumber(base)

	fmt.Printf("Number: %d\n", num)

	s := nums.ToBin(num)
	fmt.Println("As binary: " + s)
	s = nums.ToDec(num)
	fmt.Println("As decimal: " + s)
	s = nums.ToHex(num)
	fmt.Println("As hex: " + s)

	s = nums.ToBase(num, 3)
	fmt.Println("As base-3: " + s)

}
