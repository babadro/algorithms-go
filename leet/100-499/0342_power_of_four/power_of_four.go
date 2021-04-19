package _342_power_of_four

// tptl. passed. easy to remember.
func isPowerOfFour2(num int) bool {
	i := 1
	for ; i < num; i *= 4 {
	}

	return i == num
}

// passed.
func isPowerOfFour(num int) bool {
	for i := 0; i < 32; i += 2 {
		if (1 << i) == num {
			return true
		}
	}

	return false
}
