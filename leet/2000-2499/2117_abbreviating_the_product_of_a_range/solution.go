package _2117_abbreviating_the_product_of_a_range

import "strconv"

// passed. hard #ntu
// https://leetcode.com/problems/abbreviating-the-product-of-a-range/discuss/1647115/Modulo-and-Double
func abbreviateProduct(left int, right int) string {
	stuff, c, total, maxStuff := 1, 0, 0, 100000000000
	pref := 1.0

	for i := left; i <= right; i++ {
		pref *= float64(i)
		stuff *= i

		for pref >= 100000 {
			pref /= 10
			if total == 0 {
				total = 6
			} else {
				total++
			}
		}

		for stuff%10 == 0 {
			stuff /= 10
			c++
		}

		stuff %= maxStuff
	}

	s := strconv.Itoa(stuff)

	res := strconv.Itoa(int(pref))
	if total-c > 10 {
		res += "..."
	}
	if total-c >= 5 {
		res += s[len(s)-min(5, total-c-5):]
	}

	res += "e" + strconv.Itoa(c)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
