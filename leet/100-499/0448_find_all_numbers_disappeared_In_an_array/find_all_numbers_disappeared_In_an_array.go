package _448_find_all_numbers_disappeared_In_an_array

// tptl. passed
func findDisappearedNumbers(nums []int) []int {
	for i := range nums {
		for {
			idx := nums[i] - 1
			if nums[i] == nums[idx] {
				break
			}

			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}

	var res []int
	for i, num := range nums {
		if num != i+1 {
			res = append(res, i+1)
		}
	}

	return res
}
