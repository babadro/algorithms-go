package _1062_longest_repeating_substring

import "testing"

func Test_longestRepeatingSubstring(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abcd", 0},
		{"abbaba", 2},
		{"aabcaabdaab", 3},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := longestRepeatingSubstring(tt.s); got != tt.want {
				t.Errorf("longestRepeatingSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
