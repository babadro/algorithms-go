package _1770_maximum_score_from_performing_multiplication_operations

//https://leetcode.com/problems/maximum-score-from-performing-multiplication-operations/discuss/1075448/Python-DP-Clear-the-Cache!
func maximumScore2(nums []int, multipliers []int) int {
	dp := [1001][1001]int{}
	n, m := len(nums), len(multipliers)
	for l := m - 1; l >= 0; l-- {
		for r := m - 1 - l; r >= 0; r-- {
			rightIdx, mult := n-1-r, multipliers[l+r]
			dp[l][r] = max(nums[l]*mult+dp[l+1][r], nums[rightIdx]*mult+dp[l][r+1])
		}
	}

	return dp[0][0]
}

// todo 2 need to understand and fix - doesn't work
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
