package _7_minimum_deletions_to_make_sequence_sorted

import lis "github.com/babadro/algorithms-go/dp_patterns/P5_longest_common_substring/04_longest_increasing_subsequence"

func findMinDeletions(nums []int) int {
	longestSeq := lis.FindLISLenBottomUp(nums)
	return len(nums) - longestSeq
}
