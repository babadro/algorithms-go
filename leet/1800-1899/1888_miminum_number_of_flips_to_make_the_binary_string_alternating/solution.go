package _1888_miminum_number_of_flips_to_make_the_binary_string_alternating

import (
	"math"
)

// https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/discuss/1258225/Simple-cpp-solution-using-sliding-window-with-explanation.-O(N)-TC
func minFlips(s string) int {
	s += s
	length := len(s) / 2
	diff0, diff1, res := 0, 0, math.MaxInt32

	for i := 0; i < len(s); i++ {
		diff0, diff1 = diff(s, i, diff0, diff1, 1)

		if i >= length {
			diff0, diff1 = diff(s, i-length, diff0, diff1, -1)
		}

		if i >= length-1 {
			res = min3(res, diff0, diff1)
		}
	}

	return res
}

func diff(s string, i, diff0, diff1, inc int) (int, int) {
	if i%2 == 0 {
		if s[i] != '0' {
			diff0 += inc
		} else {
			diff1 += inc
		}
	} else {
		if s[i] != '1' {
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

// todo 1
func minFlips3(s string) int {
	res, s0, s1, sz := math.MaxInt32, 0, 0, len(s)
	for i := 0; i < 2*sz; i++ {
		if int(s[i%sz]) != '0'+i%2 {
			s0++
		}

		s1 += 1 - i%2
		if s[i%sz] != '0' {
			s1 += 1
		}

		if i >= sz-1 {
			if i >= sz {
				s0 -= (i - sz) % 2
				if s[i-sz] != '0' {
					s0--
				}

				s1 -= 1 - (i-sz)%2
				if s[i-sz] != '0' {
					s1--
				}
			}

			res = min3(res, s0, s1)
		}
	}

	return res
}
