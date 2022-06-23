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

// passed. top down dp
func lengthOfLIS2(nums []int) int {
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = make([]int, len(nums)+1)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	return rec(dp, nums, 0, -1)
}

func rec(dp [][]int, nums []int, cur, prev int) int {
	if cur == len(nums) {
		return 0
	}

	if dp[cur][prev+1] == -1 {
		var res1, res2 int
		if prev == -1 || nums[cur] > nums[prev] {
			res1 = 1 + rec(dp, nums, cur+1, cur)
		}

		res2 = rec(dp, nums, cur+1, prev)
		dp[cur][prev+1] = max(res1, res2)
	}

	return dp[cur][prev+1]
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
