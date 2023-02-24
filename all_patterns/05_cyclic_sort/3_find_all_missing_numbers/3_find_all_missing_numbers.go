package __find_all_missing_numbers

// leetcode 448
func findAllMissingNumbers(nums []int) []int {
	var res []int
	for i := range nums {
		for nums[i]-1 != i && nums[i] != nums[nums[i]-1] {
			j := nums[i] - 1
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	for i := range nums {
		if nums[i]-1 != i {
			res = append(res, i+1)
		}
	}

	return res
}
