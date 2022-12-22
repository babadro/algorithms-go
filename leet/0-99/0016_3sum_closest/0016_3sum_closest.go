package _016_3sum_closest

import (
	"math"
	"sort"
)

// tptl. passed. medium
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)

	minDiff := math.MaxInt64
	var sum int
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			curSum := nums[i] + nums[left] + nums[right]
			diff := curSum - target

			if diff == 0 {
				return curSum
			}

			if absDiff := abs(diff); absDiff < minDiff {
				minDiff = absDiff
				sum = curSum
			}

			if diff < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return sum
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
