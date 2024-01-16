package _2863_maximum_lengh_of_semi_decreasing_subarrays

import "sort"

// #bnsrg passed medium
// todo2 there is a solution without sorting maybe, need to check
func maxSubarrayLength(nums []int) int {
	indexes := make([]int, len(nums))

	for i := range nums {
		indexes[i] = i
	}

	sort.Slice(indexes, func(i, j int) bool {
		return nums[indexes[i]] < nums[indexes[j]]
	})

	maxIDx := -1
	res := -1

	for _, idx := range indexes {
		res = max(res, maxIDx-idx+1)

		maxIDx = max(maxIDx, idx)
	}

	return res
}
