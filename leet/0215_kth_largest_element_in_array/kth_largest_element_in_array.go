package _215_kth_largest_element_in_array

import "sort"

// TODO 1
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	pivot := nums[0]
	p1, p2 := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] >= pivot {
			p2++
		} else {
			nums[i], nums[p1] = nums[p1], nums[i]
			p1++
			p2++
		}
	}
}

// Brutforce. it works.
func findKthLargestBrutforce(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
