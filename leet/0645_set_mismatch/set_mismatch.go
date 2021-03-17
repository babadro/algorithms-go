package _645_set_mismatch

import "sort"

// todo 1 doesn't work
func findErrorNums(nums []int) []int {
	sort.Ints(nums)

	res := make([]int, 0, 2)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			res = append(res, nums[i])

			if i+1 != nums[i] {
				res = append(res, i+1)
			} else {
				res = append(res, i)
			}
		}

		if len(res) == 2 {
			return res
		}
	}

	return res
}
