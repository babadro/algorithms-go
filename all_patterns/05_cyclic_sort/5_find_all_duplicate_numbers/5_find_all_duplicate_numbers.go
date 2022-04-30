package __find_all_duplicate_numbers

// leetcode 442
func findAllDuplicates(nums []int) []int {
	for i := 0; i < len(nums); {
		j := nums[i] - 1
		if nums[i] != nums[j] {
			nums[i], nums[j] = nums[j], nums[i]
		} else {
			i++
		}
	}

	var res []int
	for i := range nums {
		if i != nums[i]-1 {
			res = append(res, nums[i])
		}
	}

	return res
}
