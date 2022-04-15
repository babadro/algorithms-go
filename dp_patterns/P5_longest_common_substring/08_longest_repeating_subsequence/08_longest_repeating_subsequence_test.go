package _8_longest_repeating_subsequence

import "testing"

func Test_findLRSLen(t *testing.T) {
	tests := []struct {
		st   string
		want int
	}{
		{"tomorrow", 2},
		{"aabdbcec", 3},
		{"fmff", 2},
	}
	for _, tt := range tests {
		t.Run(tt.st, func(t *testing.T) {
			if got := findLRSLenBottomUp(tt.st); got != tt.want {
				t.Errorf("findLRSLenBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
