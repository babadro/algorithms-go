package _747_largest_number_at_least_twice_of_others

// passed
func dominantIndex(nums []int) int {
	n := len(nums)
	if n < 2 {
		if n == 1 && nums[0] == 1 {
			return 0
		}

		return -1
	}

	max1, max2 := 0, 1
	if nums[0] < nums[1] {
		max1, max2 = 1, 0
	}

	for i := 2; i < n; i++ {
		if nums[i] > nums[max1] {
			max2 = max1
			max1 = i
		} else if nums[i] > nums[max2] {
			max2 = i
		}
	}

	if nums[max2] == 0 || nums[max1]/nums[max2] >= 2 {
		return max1
	}

	return -1
}
