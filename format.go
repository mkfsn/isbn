package isbn

import (
	"strings"
)

type Format = string

var (
	ISBN10  Format = "ISBN-10"
	ISBN13         = "ISBN-13"
	EAN13          = "EAN-13"
	Unknown        = "Unknown"
)

func rtoi(r rune) (int, error) {
	switch r {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return int(r - '0'), nil
	case 'X':
		return 10, nil
	}
	return -1, ErrInvalidISBNDigit
}

func Identify(isbn string) Format {
	if Is10(isbn) {
		return ISBN10
	} else if Is13(isbn) {
		return ISBN13
	}
	return Unknown
}

func Is10(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, c := range isbn {
		n, err := rtoi(c)
		if err != nil {
			return false
		}
		sum += (10 - i) * n
	}

	return sum%11 == 0
}

var (
	isbn13Multiplier = []int{1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1}
)

func Is13(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 13 {
		return false
	}

	sum := 0
	for i, c := range isbn {
		n, err := rtoi(c)
		if err != nil {
			return false
		}
		sum += isbn13Multiplier[i] * n
	}

	return sum%10 == 0
}
