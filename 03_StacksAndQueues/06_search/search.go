package _6_search

// leetcode 704. Binary Search
// it works.
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)
	for left < right {
		middle := (right + left) / 2
		if nums[middle] > target {
			right = middle
			continue
		} else if nums[middle] < target {
			left = middle + 1
			continue
		}
		return middle
	}
	return -1
}
