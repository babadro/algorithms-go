package _414_third_maximum_number

import "math"

// passed. tptl. easy
func thirdMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max1, max2, max3 := math.MinInt64, math.MinInt64, math.MinInt64
	for _, num := range nums {
		switch {
		case num > max1:
			max1, max2, max3 = num, max1, max2
		case num < max1 && num > max2:
			max2, max3 = num, max2
		case num < max2 && num > max3:
			max3 = num
		}
	}

	if max3 == math.MinInt64 {
		return max1
	}

	return max3
}
