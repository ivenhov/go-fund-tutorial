package main

/*
	Page 31
	Validate Luhn checksum
	- double the value of every other digit
	- add the values of the individual digits together (if a double value now has two digits, add the digits individually)
	- the identification number is valid if the sum is divisible by 10

*/
import (
	"bufio"
	"fmt"
	"os"
)

func digit(r rune) int {
	dig := r - '0'
	return int(dig)
}

func double(dig int) int {
	dig *= 2
	if dig >= 10 {
		return 1 + dig%10
	}
	return dig
}

func main() {
	in := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number: ")
	sumEven, sumOdd := 0, 0
	digs := 0
	for {
		r, _, _ := in.ReadRune()
		if r == '\n' {
			break
		}
		digs++
		dig := digit(r)
		doub := double(dig)
		if digs%2 == 0 {
			sumOdd += doub
			sumEven += dig
		} else {
			sumOdd += dig
			sumEven += doub
		}
	}
	sum := 0
	if digs%2 == 0 {
		sum = sumEven
	} else {
		sum = sumOdd
	}
	//	fmt.Printf("Luhn sum is %d\n", sum)

	if sum%10 == 0 {
		fmt.Println("Checksum valid")
	} else {
		fmt.Println("Checksum INVALID")
	}

}
