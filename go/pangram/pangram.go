// Package pangram contains functions to identify panagrams.
package pangram

import (
	"unicode"
)

// IsPangram determines whether or a string is a pangram or not
func IsPangram(s string) bool {
	if len(s) == 0 {
		return false
	}

	alphabet := make(map[rune]int, 0)
	for _, char := range s {
		if unicode.IsLetter(char) {
			alphabet[unicode.ToLower(char)]++
		}
	}

	return len(alphabet) == 26 // a - z letter count
}
