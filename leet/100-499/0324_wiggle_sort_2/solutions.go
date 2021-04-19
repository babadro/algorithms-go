package _324_wiggle_sort_2

import (
	"sort"
)

// Simple but not best
// 69% 24%
// TODO 2 look for o(n) and in place memory solution
func wiggleSort(nums []int) {
	n := len(nums)
	sort.Ints(nums)

	oddLen := n%2 != 0
	half := n / 2
	right, rightTo := n-1, half
	left, leftTo := half-1, 0
	if oddLen {
		rightTo++
		left++
		leftTo++
	}

	res := make([]int, 0, n)
	for right >= rightTo && left >= leftTo {
		res = append(res, nums[left])
		res = append(res, nums[right])
		right--
		left--
	}

	if oddLen {
		res = append(res, nums[0])
	}

	copy(nums, res)
}
