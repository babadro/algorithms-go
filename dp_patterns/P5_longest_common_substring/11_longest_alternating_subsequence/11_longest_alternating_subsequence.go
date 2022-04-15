package _1_longest_alternating_subsequence

import "github.com/babadro/algorithms-go/utils"

func findLASLenBruteForce(nums []int) int {
	res1 := recBruteForce(nums, -1, 0, true)
	res2 := recBruteForce(nums, -1, 0, false)

	return utils.Max(res1, res2)
}

func recBruteForce(nums []int, prev, cur int, isAsc bool) int {
	if cur == len(nums) {
		return 0
	}

	c1 := 0
	if isAsc {
		if prev == -1 || nums[prev] < nums[cur] {
			c1 = 1 + recBruteForce(nums, cur, cur+1, !isAsc)
		}
	} else {
		if prev == -1 || nums[prev] > nums[cur] {
			c1 = 1 + recBruteForce(nums, cur, cur+1, !isAsc)
		}
	}

	c2 := recBruteForce(nums, prev, cur+1, isAsc)

	return utils.Max(c1, c2)
}

func findLASLenTopDown(nums []int) int {
	dp := make(map[[3]int]int)
	res1 := recTopDown(dp, nums, -1, 0, true)
	res2 := recTopDown(dp, nums, -1, 0, false)

	return utils.Max(res1, res2)
}

func recTopDown(dp map[[3]int]int, nums []int, prev, cur int, isAsc bool) int {
	if cur == len(nums) {
		return 0
	}

	asc := 1
	if !isAsc {
		asc = 0
	}

	key := [3]int{prev, cur, asc}

	if _, ok := dp[key]; !ok {
		c1 := 0
		if isAsc {
			if prev == -1 || nums[prev] < nums[cur] {
				c1 = 1 + recTopDown(dp, nums, cur, cur+1, !isAsc)
			}
		} else {
			if prev == -1 || nums[prev] > nums[cur] {
				c1 = 1 + recTopDown(dp, nums, cur, cur+1, !isAsc)
			}
		}

		c2 := recTopDown(dp, nums, prev, cur+1, isAsc)

		dp[key] = utils.Max(c1, c2)
	}

	return dp[key]
}

func findLASLenBottomUp(nums []int) int {
	dp := utils.Make2DArr(len(nums), 2, 0)
	maxLen := 1
	for i := 0; i < len(nums); i++ {
		dp[i][0], dp[i][1] = 1, 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i][0] = utils.Max(dp[i][0], 1+dp[j][1])
				maxLen = utils.Max(maxLen, dp[i][0])
			} else if nums[i] != nums[j] {
				dp[i][1] = utils.Max(dp[i][1], 1+dp[j][0])
				maxLen = utils.Max(maxLen, dp[i][1])
			}
		}
	}

	return maxLen
}
