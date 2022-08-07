package _1818_min_absolute_sum_difference

import (
	"math"
)

func minAbsoluteSumDiff2(nums1 []int, nums2 []int) int {
	largestDiffIdx, maxDiff := -1, math.MinInt64
	for i := range nums1 {
		diff := abs(nums1[i] - nums2[i])

		if diff > maxDiff {
			largestDiffIdx, maxDiff = i, diff
		}
	}

	num2 := nums2[largestDiffIdx]

	idx, minDiff := -1, math.MaxInt64
	for i := range nums1 {
		diff := abs(nums1[i] - num2)
		if diff < minDiff {
			idx, minDiff = i, diff
		}
	}

	nums1[largestDiffIdx] = nums1[idx]

	res := 0
	for i := range nums1 {
		res += abs(nums1[i] - nums2[i])
		res %= 1_000_000_007
	}

	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
