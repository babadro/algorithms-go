package _131_palindrome_partitioning

import "strings"

// TODO 1
func partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	runes := []rune(s)
	n := len(runes)
	for i := 0; i < n; i++ {

	}

}

func recursive(r []rune) []rune {
	if len(r) == 0 {
		return r
	}
	return append([]rune{r[0]}, recursive(r[1:])...)
}

func isPalindrome(r []rune) bool {
	n := len(r)
	for i := 0; i < n/2; i++ {
		if r[i] != r[n-1-i] {
			return false
		}
	}
	return true
}
