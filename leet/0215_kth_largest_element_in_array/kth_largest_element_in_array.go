package _215_kth_largest_element_in_array

import "sort"

// TODO 1
func findKthLargest(nums []int, k int) int {

}

// Brutforce. it works.
func findKthLargestBrutforce(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}
