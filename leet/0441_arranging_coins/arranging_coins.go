package _441_arranging_coins

import "math"

// passed. tptl.
func arrangeCoins2(n int) int {
	l, r, last := 0, n, 0
	for l <= r {
		m := (r-l)/2 + l

		sum := seriesSum(m)
		if sum < n {
			last = m
			l = m + 1
		} else if sum > n {
			r = m - 1
		} else {
			return m
		}
	}

	return last
}

// passed. best solution if you remember approach.
func arrangeCoins3(n int) int {
	return int(math.Sqrt(2*float64(n)+0.25) - 0.5)
}

// bruteforce passed
func arrangeCoins(n int) int {
	total := n
	i := 0
	for ; i < n; i++ {
		total -= i + 1
		if total < 0 {
			return i
		}
	}

	return i
}

// sum sequence: 1+2+3+4+..+n
func seriesSum(n int) int {
	return n * (1 + n) / 2
}
