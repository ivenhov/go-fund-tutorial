package input_processing

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
)

var ILLEGAL_DIGIT = errors.New("ILLEGAL_DIGIT")

type DigitReader struct {
	r              *bufio.Reader
	SkipWhitespace bool
}

func NewDigitReader() *DigitReader {
	return &DigitReader{r: bufio.NewReader(os.Stdin)}
}

func (dr *DigitReader) ReadBinaryDigit() (int, error) {
	dig, err := dr.ReadDecimalDigit()
	if err != nil {
		return dig, err
	}
	if dig != 0 && dig != 1 {
		return 0, errors.New("Illegal digit: " + strconv.Itoa(dig))
	}
	return dig, nil
}

func (dr *DigitReader) ReadDecimalDigit() (int, error) {
	for {
		r, _, _ := dr.r.ReadRune()
		if r == '\n' {
			return 0, io.EOF
		}
		if dr.SkipWhitespace && (r == ' ' || r == '\t') {
			continue
		}
		dig, err := ToDigit(r)
		if err != nil {
			panic(err)
		}
		return dig, nil
	}
}

func (dr *DigitReader) ReadBinaryNumber() (num int) {
	dr.SkipWhitespace = true
	digs := 0
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
	return
}

func (dr *DigitReader) ReadDecimalNumber() int {
	num := 0
	for {
		r, _, _ := dr.r.ReadRune()
		if r == '\n' {
			break
		}

		dig, err := ToDigit(r)
		if err != nil {
			panic(err)
		}

		num *= 10
		num += dig
	}
	return num
}

func (dr *DigitReader) ReadHexNumber() int {
	num := 0
	for {
		r, _, _ := dr.r.ReadRune()
		if r == '\n' {
			break
		}

		dig, err := ToHexDigit(r)
		if err != nil {
			panic(err)
		}

		num *= 16
		num += dig
	}
	return num
}

func (dr *DigitReader) ReadNumber(base int) (num int) {
	for {
		r, _, _ := dr.r.ReadRune()
		if r == '\n' {
			break
		}
		dig, err := ToHexDigit(r)
		if dig >= base {
			panic("Invalid digit, base incorrect?")
		}
		if err != nil {
			panic(err)
		}

		num *= base
		num += dig
	}
	return
}
