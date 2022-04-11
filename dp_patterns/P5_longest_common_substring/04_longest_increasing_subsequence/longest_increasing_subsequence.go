package _4_longest_increasing_subsequence

import "github.com/babadro/algorithms-go/utils"

func findLISLenBruteForce(nums []int) int {
	return recBruteForce(nums, 0, -1)
}

func recBruteForce(nums []int, curr, prev int) int {
	if curr == len(nums) {
		return 0
	}

	c1 := 0
	if prev == -1 || nums[curr] > nums[prev] {
		c1 = 1 + recBruteForce(nums, curr+1, curr)
	}

	c2 := recBruteForce(nums, curr+1, prev)

	return utils.Max(c1, c2)
}

func findLISTLenTopDown(nums []int) int {
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
			c1 = 1 + recBruteForce(nums, curr+1, curr)
		}

		c2 := recBruteForce(nums, curr+1, prev)

		dp[curr][prev+1] = utils.Max(c1, c2)
	}

	return dp[curr][prev+1]
}

func FindLISLenBottomUp(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] <= dp[j]+1 {
				dp[i] = dp[j] + 1
				maxLen = utils.Max(maxLen, dp[i])
			}
		}
	}

	return dp[len(nums)-1]
}
