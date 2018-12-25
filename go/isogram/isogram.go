// Package isogram contains functions to discover words that are isograms.
package isogram

import "unicode"

// IsIsogram determines if a word is an isogram.
func IsIsogram(word string) bool {
	letterFrequency := make(map[rune]int)
	for _, char := range word {
		letter := unicode.ToLower(char)
		freq := letterFrequency[letter]
		if freq != 0 && unicode.IsLetter(letter) {
			return false
		}
		letterFrequency[letter] = 1
	}
	return true
}
