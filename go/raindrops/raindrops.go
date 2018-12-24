// Package raindrops contains functions related to number
// to string conversions.
package raindrops

import "strconv"

// Convert a number to a raindrop sound string.
func Convert(num int) string {
	var output string

	if num%3 == 0 {
		output += "Pling"
	}
	if num%5 == 0 {
		output += "Plang"
	}
	if num%7 == 0 {
		output += "Plong"
	}

	if len(output) == 0 {
		return strconv.Itoa(num)
	}

	return output
}
