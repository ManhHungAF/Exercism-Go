package diffsquares

// SquareOfSum docs
func SquareOfSum(n int) int {
	var sum int
	// for i := 1; i <= n; i++ {
	// 	sum += i
	// }
	sum = n * (n + 1) / 2
	return sum * sum
}

// SumOfSquares docs
func SumOfSquares(n int) int {
	// var sum int
	// for i := 1; i <= n; i++ {
	// 	sum += i * i
	// }
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference docs
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
