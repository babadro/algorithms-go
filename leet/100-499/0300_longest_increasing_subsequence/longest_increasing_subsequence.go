package _0300_longest_increasing_subsequence

import "sort"

// tptl passed medium (hard for me) #array #binarysearch
// todo 1 implement dp from here
// https://leetcode.com/problems/longest-increasing-subsequence/discuss/1326308
// this is an example of binary search approach
// https://leetcode.com/problems/longest-increasing-subsequence/discuss/341550/Java-solution
func lengthOfLIS(nums []int) int {
	var seq []int
	for _, num := range nums {
		n := len(seq)
		idx := sort.Search(n, func(i int) bool {
			return seq[i] >= num
		})

		if idx == n {
			seq = append(seq, num)
		} else {
			seq[idx] = num
		}
	}

	return len(seq)
}
