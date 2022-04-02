package __minimum_deletions_in_a_string_to_make_it_palindrome

import lpsq "github.com/babadro/algorithms-go/dp_patterns/P4_Palindromic_Subsequence/1_longest_palindromic_subsequence"

func findMinimumDeletions(st string) int {
	return len(st) - lpsq.FindLPSLenBottomUp(st)
}
