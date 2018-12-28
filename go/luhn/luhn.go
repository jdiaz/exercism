// Package luhn contains functions to validate credit card numbers.
package luhn

import (
	"strings"
	"unicode"
)

// Valid determines if a credit card number is valid or not.
func Valid(creditcard string) bool {
	creditcard = strings.Replace(creditcard, " ", "", -1)
	n := len(creditcard)
	if n <= 1 {
		return false
	}

	total := 0
	count := 0
	for pos := n - 1; pos >= 0; pos-- {
		c := rune(creditcard[pos])
		if !unicode.IsDigit(c) {
			return false
		}
		digit := int(c - '0')
		if count%2 != 0 { // is odd
			digit = digit * 2
			if digit > 9 {
				digit -= 9
			}
		}
		total += digit
		count++
	}

	return total%10 == 0
}
