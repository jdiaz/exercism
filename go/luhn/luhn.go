// Package luhn contains functions to validate credit card numbers.
package luhn

import (
	"strings"
	"unicode"
)

// Valid determines if a credit card number is valid or not.
func Valid(creditcard string) bool {
	creditcard = strings.Replace(creditcard, " ", "", -1)
	if !isValidCreditCardNumberString(creditcard) {
		return false
	}

	total := 0
	for pos, count := len(creditcard)-1, 0; pos >= 0; pos, count = pos-1, count+1 {
		digit := int(creditcard[pos] - '0')
		if count%2 != 0 { // is odd
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		total += digit
	}

	if total%10 == 0 {
		return true
	}

	return false
}

func isValidCreditCardNumberString(s string) bool {
	if len(s) <= 1 {
		return false
	}

	for _, char := range s {
		if !unicode.IsDigit(char) {
			return false
		}
	}
	return true
}
