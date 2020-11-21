package _231_power_of_two

// passed. tptl
func isPowerOfTwo(n int) bool {
	if n > 0 {
		for i := 0; i < 31; i++ {
			if n == 1<<i {
				return true
			}
		}
	}

	return false
}
