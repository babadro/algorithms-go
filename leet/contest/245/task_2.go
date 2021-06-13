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

	var b2 []byte
	for i := range s {
		if b[i] != 0 {
			b2 = append(b2, b[i])
		}
	}

	newS := string(b2)

	res := isSubsequence(p, newS)

	return res
}

func isSubsequence(s string, t string) bool {
	j := 0
	for i := 0; i < len(t) && j < len(s); i++ {
		if t[i] == s[j] {
			j++
		}
	}

	return j == len(s)
}
