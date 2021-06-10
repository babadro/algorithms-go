package _1888_miminum_number_of_flips_to_make_the_binary_string_alternating

import (
	"math"
)

// https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/discuss/1258225/Simple-cpp-solution-using-sliding-window-with-explanation.-O(N)-TC
// passed. tptl. medium (hard for me). #string
func minFlips(s string) int {
	n := len(s)
	diff0, diff1, res := 0, 0, math.MaxInt32

	for i := 0; i < n*2; i++ {
		diff0, diff1 = diff(s, i, diff0, diff1, 1)

		if i >= n {
			diff0, diff1 = diff(s, i-n, diff0, diff1, -1)
		}

		if i >= n-1 {
			res = min3(res, diff0, diff1)
		}
	}

	return res
}

func diff(s string, i, diff0, diff1, inc int) (int, int) {
	if i%2 == 0 {
		if s[i%len(s)] != '0' {
			diff0 += inc
		} else {
			diff1 += inc
		}
	} else {
		if s[i%len(s)] != '1' {
			diff0 += inc
		} else {
			diff1 += inc
		}
	}

	return diff0, diff1
}

func min3(a, b, c int) int {
	res := a
	if b < res {
		res = b
	}

	if c < res {
		res = c
	}

	return res
}

// todo 3 find faster solution (maybe xor)
