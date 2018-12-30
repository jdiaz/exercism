// Package wordcount contains functions for counting words in a string.
package wordcount

import (
	"strings"
	"unicode"
)

// Frequency is a map of words associated with the number of times they show up in a string.
type Frequency map[string]int

// WordCount determines the ocurrance of each word in string.
func WordCount(s string) Frequency {
	freq := Frequency{}
	buf := []rune{}
	for _, c := range strings.ToLower(s) {
		if unicode.IsLetter(c) || c == '\'' || unicode.IsDigit(c) {
			buf = append(buf, c)
			continue
		}
		// Not valid char delimiter; count word
		n := len(buf)
		if n > 0 {
			if buf[0] == '\'' && buf[n-1] == '\'' {
				// Remove single quotes
				buf = buf[1 : n-1]
			}
			freq[string(buf)]++
			buf = []rune{}
		}

	}
	// Check for last word
	if len(buf) > 0 {
		freq[string(buf)]++
	}

	return freq
}
