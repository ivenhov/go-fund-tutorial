package nums

import (
	"fmt"
)

const (
	digits = "0123456789abcdef"
)

func ToBin(num int) string {
	//	var s string
	//	mask := 1 << 31
	//	for i := 1; i <= 32; i++ {
	//		if num&mask == 0 {
	//			s += "0"
	//		} else {
	//			s += "1"
	//		}
	//		mask >>= 1
	//	}
	//	return s
	return ToBase(num, 2)
}

func _ToDec(num int) string {
	fmt.Printf("Converting to dec: %d\n", num)
	var s string
	v := num / 10
	m := num % 10
	fmt.Printf("v: %d, m: %d\n", v, m)
	if v > 0 {
		s = ToDec(v) + s
	}
	s += string(digits[m])
	fmt.Printf("Returning string %s\n", s)
	return s
}

func ToDec(num int) string {
	return ToBase(num, 10)
}

func ToHex(num int) string {
	return ToBase(num, 16)
}

func ToBase(num, base int) string {
	//	fmt.Printf("Converting to dec: %d\n", num)
	if base < 2 || base > 16 {
		panic("Unsupported base")
	}

	var b [32]byte
	i := 31
	digs := 0
	for num > 0 {
		v := num / base
		m := num % base
		//		fmt.Printf("v: %d, m: %d\n", v, m)
		b[i] = digits[m]
		i--
		digs++
		num = v
	}
	//	fmt.Printf("Bytes  %b\n", b)
	s := string(b[31-digs:])
	//	fmt.Printf("Returning string %s\n", s)
	return s
}
