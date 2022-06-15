package _442_find_all_duplicates_in_an_array

// tptl. passed
func findDuplicates(nums []int) []int {
	for i := range nums {
		for {
			idx := nums[i] - 1
			if nums[i] == nums[idx] {
				break
			}

			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}

	var res []int
	for i := range nums {
		if nums[i]-1 != i {
			res = append(res, nums[i])
		}
	}

	return res
}

// best solution
func findDuplicates2(nums []int) []int {
	var result []int
	if len(nums) == 0 {
		return result
	}

	for _, val := range nums {
		visited := abs(val) - 1
		if nums[visited] < 0 {
			result = append(result, abs(val))
		}

		nums[visited] = -nums[visited]
	}

	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}
