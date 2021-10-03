package _1955_count_number_of_special_subsequences

const mod = 1_000_000_007

// https://leetcode.com/problems/count-number-of-special-subsequences/discuss/1376530/Easy-DP-Solution-or-C%2B%2B-or-O(n)-Time-O(1)-Space-or-With-Explanation
// tptl. passed. hard. best solution. #array
func countSpecialSubsequences(nums []int) int {
	z, o, t := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			z = (2*z + 1) % mod
		} else if nums[i] == 1 {
			o = (2*o + z) % mod
		} else {
			t = (2*t + o) % mod
		}
	}

	return t
}
