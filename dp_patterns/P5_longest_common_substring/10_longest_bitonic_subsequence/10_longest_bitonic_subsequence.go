package _10_longest_bitonic_subsequence

import "github.com/babadro/algorithms-go/utils"

func findLBSLenBruteForce(nums []int) int {
	max := 0
	for i := 0; i < len(nums); i++ {
		c1 := recBruteForce(nums, i, -1)
		c2 := recBruteForceReverse(nums, i, -1)
		max = utils.Max(max, c1+c2-1)
	}

	return max
}

// find the longest decreasing subsequence from currentIndex till the end of the array
func recBruteForce(nums []int, currIDx, prevIDx int) int {
	if currIDx == len(nums) {
		return 0
	}

	c1 := 0
	if prevIDx == -1 || nums[currIDx] < nums[prevIDx] {
		c1 = 1 + recBruteForce(nums, currIDx+1, currIDx)
	}

	c2 := recBruteForce(nums, currIDx+1, prevIDx)

	return utils.Max(c1, c2)
}

// find the longest decreasing subsequence from currentIndex till the beginning of the array
func recBruteForceReverse(nums []int, currIDx, prevIDx int) int {
	if currIDx < 0 {
		return 0
	}

	c1 := 0
	if prevIDx == -1 || nums[currIDx] < nums[prevIDx] {
		c1 = 1 + recBruteForceReverse(nums, currIDx-1, currIDx)
	}

	c2 := recBruteForceReverse(nums, currIDx-1, prevIDx)

	return utils.Max(c1, c2)
}

func findLBSLenTopDown(nums []int) int {
	dp1, dp2 := makeDP(len(nums), len(nums)+1), makeDP(len(nums), len(nums)+1)
	max := 0
	for i := 0; i < len(nums); i++ {
		c1 := recTopDown(dp1, nums, i, -1)
		c2 := recTopDownReverse(dp2, nums, i, -1)

		max = utils.Max(max, c1+c2-1)
	}

	return max
}

func recTopDown(dp [][]int, nums []int, curr, prev int) int {
	if curr == len(nums) {
		return 0
	}

	if dp[curr][prev+1] == -1 {
		c1 := 0
		if prev == -1 || nums[curr] < nums[prev] {
			c1 = 1 + recTopDown(dp, nums, curr+1, curr)
		}

		c2 := recTopDown(dp, nums, curr+1, prev)

		dp[curr][prev+1] = utils.Max(c1, c2)
	}

	return dp[curr][prev+1]
}

func recTopDownReverse(dp [][]int, nums []int, curr, prev int) int {
	if curr < 0 {
		return 0
	}

	if dp[curr][prev+1] == -1 {
		c1 := 0
		if prev == -1 || nums[curr] < nums[prev] {
			c1 = 1 + recTopDownReverse(dp, nums, curr-1, curr)
		}

		c2 := recTopDownReverse(dp, nums, curr-1, prev)

		dp[curr][prev+1] = utils.Max(c1, c2)
	}

	return dp[curr][prev+1]
}

func makeDP(m, n int) [][]int {
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
		for j := range res[i] {
			res[i][j] = -1
		}
	}

	return res
}
