package _053_maxium_subarray

import "math"

// https://leetcode.com/problems/maximum-subarray/discuss/859357/Kadane's-algorithm-with-explanation
// tptl. passed. hard for me
func maxSubArray(nums []int) int {
	maxSum, curSubArraySum := math.MinInt64, 0
	for _, num := range nums {
		curSubArraySum = curSubArraySum + num
		maxSum = max(maxSum, curSubArraySum)
		if curSubArraySum < 0 {
			curSubArraySum = 0
		}
	}

	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
