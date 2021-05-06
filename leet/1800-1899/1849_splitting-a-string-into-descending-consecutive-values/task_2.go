package _1849_splitting_a_string_into_descending_consecutive_values

import "strconv"

// tptl. passed. medium
// todo 2 slightly better solution: https://leetcode.com/problems/splitting-a-string-into-descending-consecutive-values/discuss/1187821/Go
func splitString(s string) bool {
	n, i := len(s), 0
	for ; i < n-1 && s[i] == '0'; i++ {
	}

	for j := i + 1; j < n; j++ {
		num, _ := strconv.Atoi(s[i:j])
		if helper(j, s, num) {
			return true
		}
	}

	return false
}

func helper(i int, s string, last int) bool {
	n := len(s)
	if i == n {
		return true
	}

	for ; i < n-1 && s[i] == '0'; i++ {
	}

	for j := i + 1; j <= n; j++ {
		num, _ := strconv.Atoi(s[i:j])
		if num > last-1 {
			return false
		}

		if num == last-1 {
			return helper(j, s, num)
		}
	}

	return false
}
