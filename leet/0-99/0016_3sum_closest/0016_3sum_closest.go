package _016_3sum_closest

import (
	"math"
	"sort"
)

// tptl. passed. medium
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	minDiff := math.MaxInt64
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			targetDiff := target - nums[i] - nums[left] - nums[right]
			if targetDiff == 0 {
				return target
			}

			if abs(targetDiff) < abs(minDiff) {
				minDiff = targetDiff
			}

			if targetDiff > 0 {
				left++
			} else {
				right--
			}
		}
	}

	return target - minDiff
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
