package _963_minimum_number_of_swaps_to_make_the_string_balanced

// passed. tptl. medium (hard for me)
// https://leetcode.com/problems/minimum-number-of-swaps-to-make-the-string-balanced/discuss/1390576/Two-Pointers
func minSwaps(s string) int {
	res, balance := 0, 0
	for _, ch := range s {
		if ch == '[' {
			balance++
		} else {
			balance--
		}

		if balance == -1 {
			res++
			balance = 1
		}
	}

	return res
}
