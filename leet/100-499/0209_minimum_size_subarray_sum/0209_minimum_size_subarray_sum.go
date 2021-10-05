package _209_minimum_size_subarray_sum

import "math"

// tptl. passed #array #slidingwindow #medium
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	res := math.MaxInt64
	left, sum := 0, 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		for sum >= target {
			res = min(res, i+1-left)
			sum -= nums[left]
			left++
		}
	}

	if res == math.MaxInt64 {
		return 0
	}

	return res
}

// passed, but verbose.
// todo 2 check binary search solution (n * log n)
func minSubArrayLen1(target int, nums []int) int {
	curSum, n, res, curLen := 0, len(nums), math.MaxInt64, 1
	i, j := 0, 0
	for i < n {
		for ; i < n; i, curLen = i+1, curLen+1 {
			curSum += nums[i]
			if curSum >= target {
				res = min(res, curLen)
				break
			}
		}

		for ; curSum-nums[j] >= target && j < n && j < i; j++ {
			curSum -= nums[j]
			curLen--
			res = min(res, curLen)
		}

		i++
		curLen++
	}

	if res == math.MaxInt64 {
		return 0
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
