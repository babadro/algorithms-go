package _4_longest_increasing_subsequence

import "github.com/babadro/algorithms-go/utils"

func findLISLengthBruteForce(nums []int) int {
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
