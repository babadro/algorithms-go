package _442_find_all_duplicates_in_an_array

// todo 1
func findDuplicates(nums []int) []int {
	var res []int

	for i := 0; i < len(nums); i++ {
		j := abs(nums[i]) - 1
		if nums[j] >= 0 {
			nums[j] = -nums[j]
		} else {
			res = append(res, j+1)
		}
	}

	return res
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
