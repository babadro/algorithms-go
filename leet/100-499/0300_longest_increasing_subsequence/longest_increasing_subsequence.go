package _0300_longest_increasing_subsequence

// todo 1
// https://leetcode.com/problems/longest-increasing-subsequence/discuss/341550/Java-solution
func lengthOfLIS(nums []int) int {
	tails := make([]int, len(nums))
	size := 0
	for _, x := range nums {
		i, j := 0, size
		for i != j {
			m := (i + j) / 2
			if tails[m] < x {
				i = m + 1
			} else {
				j = m
			}
		}
		tails[i] = x
		if i == size {
			size++
		}
	}

	return size
}
