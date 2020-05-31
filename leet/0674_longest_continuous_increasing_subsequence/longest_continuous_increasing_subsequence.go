package _674_longest_continuous_increasing_subsequence

// ok
func findLengthOfLCIS(nums []int) int {
	maxLen, currLen := 0, 0
	for i, num := range nums {
		if currLen == 0 || num > nums[i-1] {
			currLen++
			if maxLen < currLen {
				maxLen = currLen
			}
		} else {
			currLen = 1
		}

	}
	return maxLen
}
