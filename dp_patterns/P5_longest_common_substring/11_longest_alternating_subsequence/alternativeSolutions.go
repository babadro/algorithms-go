package _1_longest_alternating_subsequence

import "github.com/babadro/algorithms-go/utils"

func findLASLenBruteForceAlternative(nums []int) int {
	return recBruteForceAlternative(nums, 0, -1, -1)
}

func recBruteForceAlternative(nums []int, cur, prev, prevPrev int) int {
	if cur == len(nums) {
		return 0
	}

	c1 := 0
	if prev == -1 || (prevPrev == -1 && nums[cur] != nums[prev]) ||
		(nums[cur] > nums[prev] && nums[prev] < nums[prevPrev]) ||
		(nums[cur] < nums[prev] && nums[prev] > nums[prevPrev]) {
		c1 = 1 + recBruteForceAlternative(nums, cur+1, cur, prev)
	}

	c2 := recBruteForceAlternative(nums, cur+1, prev, prevPrev)

	return utils.Max(c1, c2)
}

// doesn't work. as slow as bruteforce solution
func findLasLenTopDownAlternative(nums []int) int {
	dp := make(map[[3]int]int)
	return recTopDownAlternative(dp, nums, 0, -1, -1)
}

func recTopDownAlternative(dp map[[3]int]int, nums []int, cur, prev, prevPrev int) int {
	if cur == len(nums) {
		return 0
	}

	key := [3]int{cur, prev, prevPrev}
	if _, ok := dp[key]; !ok {
		c1 := 0
		if prev == -1 || (prevPrev == -1 && nums[cur] != nums[prev]) ||
			(nums[cur] > nums[prev] && nums[prev] < nums[prevPrev]) ||
			(nums[cur] < nums[prev] && nums[prev] > nums[prevPrev]) {
			c1 = 1 + recBruteForceAlternative(nums, cur+1, cur, prev)
		}

		c2 := recBruteForceAlternative(nums, cur+1, prev, prevPrev)

		dp[key] = utils.Max(c1, c2)
	}

	return dp[key]
}

func findLASLenTopDown2DP(nums []int) int {
	dp1 := utils.Make2DArr(len(nums)+1, len(nums), -1)
	dp2 := utils.Make2DArr(len(nums)+1, len(nums), -1)

	res1 := recTopDown2DP(dp1, nums, -1, 0, true)
	res2 := recTopDown2DP(dp2, nums, -1, 0, false)

	return utils.Max(res1, res2)
}

func recTopDown2DP(dp [][]int, nums []int, prev, cur int, isAsc bool) int {
	if cur == len(nums) {
		return 0
	}

	if dp[prev+1][cur] == -1 {
		c1 := 0
		if isAsc {
			if prev == -1 || nums[prev] < nums[cur] {
				c1 = 1 + recTopDown2DP(dp, nums, cur, cur+1, !isAsc)
			}
		} else {
			if prev == -1 || nums[prev] > nums[cur] {
				c1 = 1 + recTopDown2DP(dp, nums, cur, cur+1, !isAsc)
			}
		}

		c2 := recTopDown2DP(dp, nums, prev, cur+1, isAsc)

		dp[prev+1][cur] = utils.Max(c1, c2)
	}

	return dp[prev+1][cur]
}

// x4 times faster than topDown with map, but x4 slower, than findLASLenTopDown2DP
func findLASLenTopDown3DArr(nums []int) int {
	dp := utils.Make3DArr(len(nums)+1, len(nums), 2, -1)
	res1 := recTopDown3DArr(dp, nums, -1, 0, true)
	res2 := recTopDown3DArr(dp, nums, -1, 0, false)

	return utils.Max(res1, res2)
}

func recTopDown3DArr(dp [][][]int, nums []int, prev, cur int, isAsc bool) int {
	if cur == len(nums) {
		return 0
	}

	asc := 1
	if !isAsc {
		asc = 0
	}

	if dp[prev+1][cur][asc] == -1 {
		c1 := 0
		if isAsc {
			if prev == -1 || nums[prev] < nums[cur] {
				c1 = 1 + recTopDown3DArr(dp, nums, cur, cur+1, !isAsc)
			}
		} else {
			if prev == -1 || nums[prev] > nums[cur] {
				c1 = 1 + recTopDown3DArr(dp, nums, cur, cur+1, !isAsc)
			}
		}

		c2 := recTopDown3DArr(dp, nums, prev, cur+1, isAsc)

		dp[prev+1][cur][asc] = utils.Max(c1, c2)
	}

	return dp[prev+1][cur][asc]
}
