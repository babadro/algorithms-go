package _326_power_of_three

// TODO 1
func isPowerOfThreeNaive(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	for {
		if n%3 != 0 {
			return false
		}
		n = n / 3
		if n == 1 {
			return true
		}
	}
}
