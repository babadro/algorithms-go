package _442_find_all_duplicates_in_an_array

// tptl. passed
func findDuplicates(nums []int) []int {
	for i := 0; i < len(nums); {
		j := nums[i] - 1
		if i == j || nums[i] == nums[j] {
			i++
			continue
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	var res []int
	for i := range nums {
		if i != nums[i]-1 {
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
