// Package diffsquares implements difference of squares algorithm
package diffsquares

// SumOfSquares returns the sum of the squares of the first n natural numbers
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// SquareOfSum returns the square of the sum of the first n natural numbers
func SquareOfSum(n int) int {
	sum := n * (n + 1) / 2
	return sum * sum
}

// Difference returns the difference between the sum of squares and the square of sum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

