package _3_minimum_deletions_and_insertions_to_transform_string_to_another

import subSeq "github.com/babadro/algorithms-go/dp_patterns/P5_longest_common_substring/02_longest_common_subsequence"

// minimum deletions or insertions to convert s1 to s2
func findMDI(s1, s2 string) int {
	lSubSeq := subSeq.FindLCSLenBottomUp(s1, s2)

	// delete everything from s1 which is not part of lSubSeq
	// insert everything in s1 which is present in s2 but not part of lSubSeq,
	return len(s1) - lSubSeq + len(s2) - lSubSeq
}
