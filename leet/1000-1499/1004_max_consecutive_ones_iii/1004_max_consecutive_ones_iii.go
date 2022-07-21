package _1004_max_consecutive_ones_iii

// tptl passed medium. best solution
func longestOnes(nums []int, k int) int {
	left, res, zeroes := 0, 0, 0
	for right := range nums {
		if nums[right] == 0 {
			zeroes++
		}

		if zeroes > k {
			if nums[left] == 0 {
				zeroes--
			}

			left++
		}

		res = max(res, right-left+1)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
