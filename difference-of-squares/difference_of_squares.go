package diffsquares

// SumOfSquares solution
func SumOfSquares(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}


// SquareOfSum solution
func SquareOfSum(n int) int {
	sum := (1 + n) * n / 2
	return sum * sum
}

// Difference solution
func Difference(n int) int {
	var sum int
	if n <= 1 {
		return 0
	}
	for i := 1; i <= n; i++ {
		sum += i * i * (i - 1)
	}
	return sum
}