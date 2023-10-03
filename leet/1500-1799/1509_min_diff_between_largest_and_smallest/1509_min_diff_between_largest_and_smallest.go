package _1509_min_diff_between_largest_and_smallest

import (
	"math"
)

// bnsrg #medium. passed, fast
func minDifference(nums []int) int {
	if len(nums) < 5 {
		return 0
	}

	max1, max2, max3, max4 := math.MinInt64, math.MinInt64, math.MinInt64, math.MinInt64
	min1, min2, min3, min4 := math.MaxInt64, math.MaxInt64, math.MaxInt64, math.MaxInt64

	for _, num := range nums {
		switch {
		case num > max1:
			max1, max2, max3, max4 = num, max1, max2, max3
		case num > max2:
			max2, max3, max4 = num, max2, max3
		case num > max3:
			max3, max4 = num, max3
		case num > max4:
			max4 = num
		}

		switch {
		case num < min1:
			min1, min2, min3, min4 = num, min1, min2, min3
		case num < min2:
			min2, min3, min4 = num, min2, min3
		case num < min3:
			min3, min4 = num, min3
		case num < min4:
			min4 = num
		}
	}

	return min(max1-min4, max4-min1, max2-min3, max3-min2)
}
