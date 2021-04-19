package _459_repeated_substring_pattern

import "testing"

func Test_repeatedSubstringPattern(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"abab", true},
		{"aba", false},
		{"abaaba", true},
		{"aaaa", true},
		{"abcabcabcabc", true},
		{"abcabcabcabc1", false},
		{"a", false},
		{"babbabbabbabbab", true},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := repeatedSubstringPattern2(tt.s); got != tt.want {
				t.Errorf("repeatedSubstringPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
