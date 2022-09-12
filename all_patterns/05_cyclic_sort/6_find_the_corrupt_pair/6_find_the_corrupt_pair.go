package __find_the_corrupt_pair

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
