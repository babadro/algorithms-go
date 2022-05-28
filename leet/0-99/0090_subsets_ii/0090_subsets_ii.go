package _090_subsets_ii

import "sort"

// tptl. passed
func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	res := [][]int{{}}
	prevEnd := 0
	for i := range nums {
		start := 0
		if i > 0 && nums[i] == nums[i-1] {
			start = prevEnd + 1
		}

		prevEnd = len(res) - 1
		for j := start; j <= prevEnd; j++ {
			arr := make([]int, len(res[j])+1)
			copy(arr, res[j])
			arr[len(arr)-1] = nums[i]

			res = append(res, arr)
		}
	}

	return res
}
