package _1888_miminum_number_of_flips_to_make_the_binary_string_alternating

import (
	"github.com/babadro/algorithms-go/utils"
	"math"
)

//https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/discuss/1253967/C%2B%2B-O(N)-Time-or-O(1)-Space-or-Intuitive
// https://leetcode.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/discuss/1258225/Simple-cpp-solution-using-sliding-window-with-explanation.-O(N)-TC

// todo 1
func minFlips(s string) int {
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

			res = utils.Min3(res, s0, s1)
		}
	}

	return res
}
