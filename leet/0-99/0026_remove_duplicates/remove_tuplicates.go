package _026_remove_duplicates

// tptl.
func removeDuplicates(nums []int) int {
	i := 1
	for j := 1; j < len(nums); j++ {
		if nums[j] == nums[i-1] {
			continue
		}

		nums[i] = nums[j]
		i++
	}

	return i
}
