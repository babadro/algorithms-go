package _033_search_in_rotated_sorted_array

import (
	"sort"
)

// TODO 2 look for mor short solution. And, maybe, without sort.Search
// like https://leetcode.com/problems/search-in-rotated-sorted-array/discuss/307436/Faster-than-100-Go
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Search in Rotated Sorted Array.
// Memory Usage: 2.6 MB, less than 50.00% of Go online submissions for Search in Rotated Sorted Array.
func search(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return -1
	}
	if l == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	leftIdx, rightIdx := 0, l-1
	for rightIdx-leftIdx > 1 {
		idx := ((rightIdx - leftIdx) / 2) + leftIdx
		if nums[idx] < nums[leftIdx] {
			rightIdx = idx
		} else {
			leftIdx = idx
		}
	}
	if sliceIsSorted := nums[rightIdx] > nums[leftIdx]; sliceIsSorted {
		return binarySearch(nums, target, 0)
	}
	if target <= nums[l-1] {
		return binarySearch(nums[rightIdx:], target, rightIdx)
	}
	return binarySearch(nums[:rightIdx], target, 0)
}

func binarySearch(sliceOfNums []int, target int, offset int) int {
	idx := sort.Search(len(sliceOfNums), func(i int) bool {
		return sliceOfNums[i] >= target
	})
	if idx < len(sliceOfNums) && sliceOfNums[idx] == target {
		return idx + offset
	}
	return -1
}
