// Package proverb contains functions for handling proverbs.
package proverb

import "fmt"

// Proverb returns an array of verses that make up a proverb given a list of
// rhymes.
func Proverb(rhyme []string) []string {
	n := len(rhyme)
	if n == 0 {
		return rhyme
	}
	proverb := []string{}
	for pos := 1; pos < n; pos++ {
		proverb = append(proverb, fmt.Sprintf("For want of a %s the %s was lost.", rhyme[pos-1], rhyme[pos]))
	}
	return append(proverb, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
}
