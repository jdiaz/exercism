// Package diffsquares contains functions to calculate SquareOfSum, SumOfSquares, and their difference.
package diffsquares

import "math"

// SquareOfSum computes the square of the sum of the first n numbers.
func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

// SquareOfSum computes the sum of squares of the first n numbers.
func SumOfSquares(n int) int {
	fn := float64(n)
	total := 0.0
	for i := 1.0; i <= fn; i++ {
		total += math.Pow(i, 2)
	}
	return int(total)
}

// Difference computes the difference between SquareOfSum and SumOfSquares values.
func Difference(n int) int {
	return int(SquareOfSum(n) - SumOfSquares(n))
}
