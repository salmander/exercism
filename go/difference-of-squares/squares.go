package diffsquares

func SquareOfSums(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += i
	}
	return sum * sum
}

func SumOfSquares(limit int) int {
	sum := 0
	for i := 1; i <= limit; i++ {
		sum += i * i
	}
	return sum
}

func Difference(number int) int {
	return SquareOfSums(number) - SumOfSquares(number)
}
