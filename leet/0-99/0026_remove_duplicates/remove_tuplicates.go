package _026_remove_duplicates

import (
	"math"
	"sort"
)

// tptl. #hard for me
func removeDuplicates(nums []int) int {
	lastPos := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[lastPos] {
			nums[lastPos+1] = nums[i]
			lastPos++
		}
	}

	return lastPos + 1
}

func removeDuplicates2(nums []int) int {
	last := 0
	counter := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[last] {
			nums[i] = math.MaxInt64
			counter++
		} else {
			last = i
		}
	}
	sort.Ints(nums)
	return len(nums) - counter
}
