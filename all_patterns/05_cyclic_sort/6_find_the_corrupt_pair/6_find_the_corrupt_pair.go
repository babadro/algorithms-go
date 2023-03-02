package __find_the_corrupt_pair

// Given a non-empty unsorted array taken from a range of 1 to n.
// Due to some data error, one of the numbers is duplicated,
// which results in another number missing.
// Create a function that returns the corrupt pair (missing, duplicated).
func findCorruptNums(nums []int) (int, int) {
	for i := range nums {
		for {
			idx := nums[i] - 1
			if nums[idx] == nums[i] {
				break
			}

			nums[idx], nums[i] = nums[i], nums[idx]
		}
	}

	for i := range nums {
		if nums[i]-1 != i {
			return nums[i], i + 1
		}
	}

	return -1, -1
}
