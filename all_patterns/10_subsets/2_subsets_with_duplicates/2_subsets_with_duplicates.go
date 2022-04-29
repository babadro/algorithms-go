package __subsets_with_duplicates

// todo 1 similar problem - solve this task in O(1) space without modifying input using two pointers
func findNumber(nums []int) int {
	for i := 0; i < len(nums); {
		j := nums[i] - 1
		if j == i {
			i++
			continue
		}

		if nums[i] == nums[j] {
			return nums[i]
		}

		nums[i], nums[j] = nums[j], nums[i]
	}

	return 0
}
