package input_processing

import (
	"errors"
)

func ToDigit(r rune) (int, error) {
	if r >= '0' && r <= '9' {
		dig := r - '0'
		return int(dig), nil
	} else {
		return -1, errors.New("Not a digit '" + string(r) + "'")
	}
}

func ToHexDigit(r rune) (int, error) {
	if r >= '0' && r <= '9' {
		return int(r - '0'), nil
	}
	if r >= 'a' && r <= 'f' {
		return int(r - 'a' + 10), nil
	}
	if r >= 'A' && r <= 'F' {
		return int(r - 'A' + 10), nil
	}
	return -1, errors.New("Not a digit '" + string(r) + "'")
}
