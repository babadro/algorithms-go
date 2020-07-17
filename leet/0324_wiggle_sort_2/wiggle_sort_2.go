package _324_wiggle_sort_2

import "sort"

// Doesn't work
func wiggleSortBad(nums []int) {
	n := len(nums)
	sort.Ints(nums)
	for i := 0; i < n; i += 2 {
		if i > 0 && nums[i] > nums[i-1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
		if i+1 == n {
			break
		}
		if nums[i+1] == nums[i] {
			j := i + 2
			for nums[j] == nums[i+1] {
				j++
			}
			nums[i+1], nums[j] = nums[j], nums[i+1]
		}
	}
}
