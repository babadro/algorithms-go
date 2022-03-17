package _2167_minimum_time_to_remove_all_cars_containing_illegal_goods

// tptl passed #hard #ntu
// https://leetcode.com/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/discuss/1748704/JavaC%2B%2BPython-Short-One-Pass-O(1)-Space
// todo 2 implement dp solution for fun
func minimumTime(s string) int {
	n, left, res := len(s), 0, len(s)

	for i := 0; i < n; i++ {
		left = min(left+int(s[i]-'0')*2, i+1)
		res = min(res, left+n-1-i)
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
