package _5_maximum_sum_increasing_subsequence

import "github.com/babadro/algorithms-go/utils"

// leetcode 1626? https://leetcode.com/problems/best-team-with-no-conflicts/discuss/899452/C%2B%2B-DP-with-comments-Find-the-maximum-sum-of-increasing-subsequence
// leetcode 300?
func findMSISBruteForce(nums []int) int {
	return recBruteForce(nums, 0, -1)
}

func recBruteForce(nums []int, curr, prev int) int {
	if curr == len(nums) {
		return 0
	}

	c1 := 0
	if prev == -1 || nums[curr] > nums[prev] {
		c1 = nums[curr] + recBruteForce(nums, curr+1, curr)
	}

	c2 := recBruteForce(nums, curr+1, prev)

	return utils.Max(c1, c2)
}

func findMSISTopDown(nums []int) int {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return recTopDown(dp, nums, 0, -1)
}

func recTopDown(dp [][]int, nums []int, curr, prev int) int {
	if curr == len(nums) {
		return 0
	}

	if dp[curr][prev+1] == -1 {
		c1 := 0
		if prev == -1 || nums[curr] > nums[prev] {
			c1 = nums[curr] + recBruteForce(nums, curr+1, curr)
		}

		c2 := recBruteForce(nums, curr+1, prev)

		dp[curr][prev+1] = utils.Max(c1, c2)
	}

	return dp[curr][prev+1]
}

func findMSISBottomUp(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = nums[i]
	}

	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+nums[i] {
				dp[i] = dp[j] + nums[i]
				maxSum = utils.Max(maxSum, dp[i])
			}
		}
	}

	return maxSum
}
