package _1_palindrome_permutation

import "testing"

func Test_permutePalindrome(t *testing.T) {
	tests := []struct {
		st   string
		want bool
	}{
		{"a", true},
		{"ab", false},
		{"", true},
		{"aaa", true},
		{"aaab", false},
	}
	for _, tt := range tests {
		t.Run(tt.st, func(t *testing.T) {
			if got := permutePalindrome(tt.st); got != tt.want {
				t.Errorf("permutePalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
