package _0300_longest_increasing_subsequence

import "sort"

// tptl passed medium (hard for me) #array #binarysearch
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

// passed. dp variant. tptl
// https://leetcode.com/problems/longest-increasing-subsequence/discuss/1326308
func lengthOfLISDP(nums []int) int {
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := range nums {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[i] < dp[j]+1 {
				dp[i] = dp[j] + 1
			}
		}
	}

	max := 0
	for _, el := range dp {
		if el > max {
			max = el
		}
	}

	return max
}
