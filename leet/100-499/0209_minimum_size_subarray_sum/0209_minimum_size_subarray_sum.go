package _209_minimum_size_subarray_sum

import "math"

// todo 1
func minSubArrayLen(target int, nums []int) int {
	curSum, n, res, curLen := 0, len(nums), math.MaxInt64, 0
	i, j := 0, 0
	for ; i < n; i, curLen = i+1, curLen+1 {
		curSum += nums[i]
		if curSum >= target {
			res = min(res, curSum)
			break
		}
	}

	for ; j < n && j < i; j, curLen = j+1, curLen-1 {
		curSum -= nums[i]
		if curSum
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
