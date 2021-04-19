package _029_divide_two_integers

import "math"

// TODO 2 This is brute force (37% cpu and 100% memory). Implement best solution
func divide(dividend int, divisor int) int {
	res := 0
	var signed bool
	if divisor == 1 {
		return checkOverflow(dividend)
	} else if divisor == -1 {
		return checkOverflow(-dividend)
	}
	if dividend < 0 {
		dividend *= -1
		if divisor < 0 {
			divisor *= -1
		} else {
			signed = true
		}
	} else if divisor < 0 {
		divisor *= -1
		signed = true
	}
	for dividend >= divisor {
		dividend -= divisor
		res += 1
	}
	if signed {
		return -res
	}
	return checkOverflow(res)
}

func checkOverflow(input int) int {
	if input > math.MaxInt32 {
		return math.MaxInt32
	} else if input < math.MinInt32 {
		return math.MinInt32
	}
	return input
}
