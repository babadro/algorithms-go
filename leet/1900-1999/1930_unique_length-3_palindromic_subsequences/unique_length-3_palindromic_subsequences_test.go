package _1930_unique_length_3_palindromic_subsequences

import "testing"

func Test_countPalindromicSubsequence(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"aabca", 3},
		{"abc", 0},
		{"bbcbaba", 4},
		{"aba", 1},
		{"aaaaaa", 1},
		{"aabbaa", 2},
		{tleInput1, 676},
		{tleInput2, 1},
		{bigInput3, 26},
	}
	for _, tt := range tests {
		name := tt.s
		if len(tt.s) > 10 {
			name = tt.s[:10]
		}
		t.Run(name, func(t *testing.T) {
			if got := countPalindromicSubsequence(tt.s); got != tt.want {
				t.Errorf("countPalindromicSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
