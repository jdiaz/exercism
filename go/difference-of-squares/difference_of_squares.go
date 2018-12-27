// Package diffsquares contains functions to calculate SquareOfSum, SumOfSquares, and their difference.
package diffsquares

// SquareOfSum computes the square of the sum of the first n numbers.
func SquareOfSum(n int) int {
	sum := (n * (n + 1)) / 2
	return sum * sum
}

// SumOfSquares computes the sum of squares of the first n numbers.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1) / 6)
}

// Difference computes the difference between SquareOfSum and SumOfSquares values.
func Difference(n int) int {
	return int(SquareOfSum(n) - SumOfSquares(n))
}
