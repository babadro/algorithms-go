package _459_repeated_substring_pattern

import "strings"

//  best solution. tptl
// todo 3 find the explanation.
// some explanation is here, but it's not clear: https://leetcode.com/problems/repeated-substring-pattern/discuss/872901/Java-One-Liner-solution-oror-Full-explanation

// -- String A is made of repeated substrings, if and only if it is a non trivial rotation of itself.
// -- If String 'A' and 'B' have the same length then we can say that Srtring 'A' is rotation of B if and only if 'A' is Substring of 'BB'.
// let's say A=reshna
// B=naresh
// then BB=nareshnaresh and we can say that A is substring of BB
// then A is confirmed rotation of B.
// -- E.g. str=abcab
// s1=str+str=abcab
// Now check the first occurance of str in s1 from index 1 to length of s1.
// Let's say we find the index is k. Then if k==length(str) then it is not valid return false
// otherwise , if it exist without any confliction then there is a solution
func repeatedSubstringPattern2(s string) bool {
	return strings.Contains((s + s)[1:2*len(s)-1], s)
}

// passed
func repeatedSubstringPattern(s string) bool {
Loop:
	for n := 1; n <= len(s)/2; n++ {
		if len(s)%n != 0 {
			continue Loop
		}

		subStr := s[:n]
		for j := 0; j <= len(s)-len(subStr); j += len(subStr) {
			if s[j:j+len(subStr)] != subStr {
				continue Loop
			}
		}

		return true
	}

	return false
}
