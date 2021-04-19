package _326_power_of_three

func isPowerOfThree(n int) bool {
	return n > 0 && 4052555153018976267%n == 0
}

func isPowerOfThreeNaive(n int) bool {
	if n < 1 {
		return false
	}
	for n%3 == 0 {
		n = n / 3
	}
	return n == 1
}
