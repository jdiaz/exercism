// Package hamming provides distance calculations.
package hamming

import "errors"

// Distance calculation of the hamming distance for two DNA string encoded sequences.
func Distance(a, b string) (int, error) {
	aSize := len(a)
	bSize := len(b)
	if aSize != bSize {
		return 0, errors.New("DNA sequence of unequal length provided")
	}
	distance := 0
	for i := 0; i < aSize; i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
