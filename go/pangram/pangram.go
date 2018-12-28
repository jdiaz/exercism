// Package pangram contains functions to identify panagrams.
package pangram

import (
	"unicode"
)

// IsPangram determines if a string is a pangram.
func IsPangram(s string) bool {
	alphabet := make(map[rune]int, 0)
	for _, char := range s {
		if unicode.IsLetter(char) {
			alphabet[unicode.ToLower(char)]++
		}
	}

	return len(alphabet) == 26 // a - z letter count
}
