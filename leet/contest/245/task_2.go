package _45

// todo 1 refac and understand
// https://leetcode.com/problems/maximum-number-of-removable-characters/discuss/1268656/C%2B%2B-Binary-search
func maximumRemovals(s string, p string, removable []int) int {
	left, right := 0, len(removable)+1
	var res int
	for left < right {
		mid := (left + right) / 2
		if check(s, p, mid, removable) {
			res = mid
			left = mid + 1
		} else {
			right = mid
		}
	}

	return res
}

func check(s, p string, count int, removable []int) bool {
	var b = []byte(s)

	for i := 0; i < count; i++ {
		b[removable[i]] = 0
	}

	j := 0
	for i := 0; i < len(b) && j < len(p); i++ {
		if p[j] == b[i] {
			j++
		}
	}

	return j == len(p)
}
