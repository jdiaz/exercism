// Package grains contains functions to calculate the number of grains of wheat on a chessboard.
package grains

import (
	"errors"
	"math"
)

const maxNumSquare = 64

// Square calculate the number of grains of wheat on a particular square of a chessboard.
func Square(num int) (uint64, error) {
	if num <= 0 || num > maxNumSquare {
		return 0, errors.New("Square input outside valid range 64 > num >= 0")
	}

	return uint64(math.Pow(2, float64(num-1))), nil
}

// Total returns the number of grains on the board
func Total() uint64 {
	var total uint64
	for i := 1; i <= maxNumSquare; i++ {
		grains, _ := Square(i)
		total += grains
	}

	return total
}
