package _041_first_missing_positive

// tptl. passed. hard
func firstMissingPositive(nums []int) int {
	for i := range nums {
		for nums[i] > 0 && nums[i] < len(nums) {
			idx := nums[i] - 1
			if nums[i] == nums[idx] {
				break
			}

			nums[i], nums[idx] = nums[idx], nums[i]
		}
	}

	for i, num := range nums {
		if num != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}
