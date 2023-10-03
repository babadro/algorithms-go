package _1509_min_diff_between_largest_and_smallest

import (
	"math"
)

func minDifference(nums []int) int {
	if len(nums) < 5 {
		return 0
	}

	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	min1, min2, min3 := math.MaxInt64, math.MaxInt64, math.MaxInt64

	for _, num := range nums {
		switch {
		case num > max1:
			max1, max2, max3 = num, max1, max2
		case num > max2:
			max2, max3 = num, max2
		case num > max3:
			max3 = num
		}

		switch {
		case num < min1:
			min1, min2, min3 = num, min1, min2
		case num < min2:
			min2, min3 = num, min2
		case num < min3:
			min3 = num
		}
	}

	return min(max2-min1, max1-min2)
}
