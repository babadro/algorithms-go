package _029_divide_two_integers

import "math"

// tptl. passed
// https://leetcode.com/problems/divide-two-integers/discuss/1007005/Simple-code-simple-explanation
func divide(dividend int, divisor int) int {
	negative := false
	if (dividend < 0 && divisor > 0) ||
		(dividend > 0 && divisor < 0) {
		negative = true
	}

	dividend, divisor = abs(dividend), abs(divisor)
	var res int
	for dividend >= divisor {
		tempDivisor, count := divisor, 1
		for dividend >= tempDivisor<<1 {
			tempDivisor <<= 1
			count <<= 1
		}

		dividend -= tempDivisor
		res += count
	}

	if negative {
		return -res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
