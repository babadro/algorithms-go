package _215_kth_largest_element_in_array

import "sort"

// Brutforce. it works.
func findKthLargestBrutforce(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
