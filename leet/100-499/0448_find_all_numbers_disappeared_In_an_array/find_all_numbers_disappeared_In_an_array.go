package _448_find_all_numbers_disappeared_In_an_array

// tptl. passed
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); {
		idx := nums[i] - 1
		if nums[i] == nums[idx] {
			i++
			continue
		}

		nums[idx], nums[i] = nums[i], nums[idx]
	}

	res := []int{}
	for i, num := range nums {
		if i+1 != num {
			res = append(res, i+1)
		}
	}

	return res
}
