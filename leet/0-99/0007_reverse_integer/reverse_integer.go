package _007_reverse_integer

import "math"

// passed. tptl
func reverse(x int) int {
	res := 0
	for x != 0 {
		res *= 10
		res += x % 10
		x /= 10

		if res > math.MaxInt32 || res < math.MinInt32 {
			return 0
		}
	}

	return res
}
