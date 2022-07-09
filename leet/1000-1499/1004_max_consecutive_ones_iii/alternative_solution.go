package _1004_max_consecutive_ones_iii

// it is slower due inner loop
func longestOnes2(nums []int, k int) int {
	zeroes, left, maxLen := 0, 0, 0
	for right := range nums {
		if nums[right] == 0 {
			zeroes++
		}

		for zeroes > k {
			if nums[left] == 0 {
				zeroes--
			}

			left++
		}

		maxLen = max(maxLen, right-left+1)
	}

	return maxLen
}
