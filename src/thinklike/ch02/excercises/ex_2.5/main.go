package main

/*
	ISBN-13 validation
*/
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"thinklike/ch02/input_processing"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter ISBN-13: ")

	digsNum := 0
	sum := 0
	exp := -1
	for {
		r, _, _ := in.ReadRune()
		if r == '\n' {
			break
		}
		if r == '-' {
			continue
		}

		dig, err := input_processing.ToDigit(r)
		if err != nil {
			panic(err)
		}
		digsNum++

		switch digsNum {
		case 1, 3, 5, 7, 9, 11:
			sum += dig
		case 2, 4, 6, 8, 10, 12:
			sum += 3 * dig
		case 13:
			exp = dig
		}
	}
	if digsNum != 13 {
		panic("Invalid number of digits: " + strconv.Itoa(digsNum) + ". Expected 13")
	}

	dif := (10 - sum%10) % 10
	if dif == exp {
		fmt.Printf("Checksum valid: %d\n", exp)
	} else {
		fmt.Printf("Checksum INVALID, wanted: %d got: %d\n", dif, exp)
	}

}
