package _1004_max_consecutive_ones_iii

// tptl passed medium. best solution
func longestOnes(nums []int, k int) int {
	start, ones, res := 0, 0, 0
	for end := range nums {
		if nums[end] == 1 {
			ones++
		}

		if end-start+1-ones > k {
			if nums[start] == 1 {
				ones--
			}

			start++
		} else {
			res = max(res, end-start+1)
		}

	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
