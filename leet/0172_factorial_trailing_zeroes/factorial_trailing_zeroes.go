package _172_factorial_trailing_zeroes

// TODO 3 I don't understand
func trailingZeroes(n int) int {
	res := 0
	for n != 0 {
		res += n / 5
		n /= 5
	}
	return res
}

func trailingZeroes2(n int) int {
	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}
	counter := 0
	for {
		if factorial%10 == 0 {
			counter++
			factorial = factorial / 10
		} else {
			break
		}
	}
	return counter
}
