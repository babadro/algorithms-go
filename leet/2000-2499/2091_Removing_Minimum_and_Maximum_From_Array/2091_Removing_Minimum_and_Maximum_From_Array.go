package Removing_Minimum_and_Maximum_From_Array

import (
	"math"
)

// passed. easy.
func minimumDeletions(nums []int) int {
	minNum, maxNum := math.MaxInt64, math.MinInt64
	var minIdx, maxIdx int
	for i, num := range nums {
		if num > maxNum {
			maxNum = num
			maxIdx = i
		}

		if num < minNum {
			minNum = num
			minIdx = i
		}
	}

	left, right := minIdx, maxIdx
	if left > right {
		left, right = right, left
	}

	res1 := right + 1
	res2 := len(nums) - left
	res3 := left + 1 + len(nums) - right

	return min3(res1, res2, res3)
}

func min3(a, b, c int) int {
	res := a
	if b < res {
		res = b
	}

	if c < res {
		res = c
	}

	return res
}
