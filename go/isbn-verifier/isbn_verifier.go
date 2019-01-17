// Package isbn contains functions for facilitating isbn manipulations.
package isbn

import (
	"strconv"
	"strings"
	"unicode"
)

const checkCharacter = 'X'

// IsValidISBN determines if a isbn number is valid.
func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	n := len(isbn)
	if isbn == "" || n != 10 {
		return false
	}
	acc := 0
	var digit int
	for index, char := range isbn {
		if char == checkCharacter && index != n-1 {
			return false
		}
		if char != checkCharacter && !unicode.IsDigit(char) {
			return false
		}
		if char != checkCharacter {
			number, _ := strconv.ParseInt(string(char), 10, 32)
			digit = int(number)
		} else {
			digit = 10
		}

		acc += (10 - index) * digit
	}

	return acc%11 == 0
}
