package _1770_maximum_score_from_performing_multiplication_operations

// todo1 need to understand
//https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/discuss/1075496/C%2B%2B-Classic-
func maximumScore(nums []int, multipliers []int) int {
	dp := [1001][1001]int{}
	return dfs(nums, multipliers, 0, 0, &dp)
}

func dfs(nums, mults []int, l, i int, dp *[1001][1001]int) int {
	if i >= len(mults) {
		return 0
	}

	if dp[l][i] == 0 {
		r := len(nums) - 1 - i - l
		dp[l][i] = max(nums[l]*mults[i]+dfs(nums, mults, l+1, i+1, dp),
			nums[r]*mults[i]+dfs(nums, mults, l, i+1, dp))
	}

	return dp[l][i]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
