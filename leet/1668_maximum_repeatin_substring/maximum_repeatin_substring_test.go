package _1668_maximum_repeatin_substring

import "testing"

func Test_maxRepeating(t *testing.T) {
	tests := []struct {
		sequence string
		word     string
		want     int
	}{
		{"ababc", "ab", 2},
		{"ababc", "ba", 1},
		{"ababc", "ac", 0},
		{"ababcabababa", "ab", 3},
		{"ababab", "ab", 3},
		{"aaaa", "a", 4},
		{"a", "a", 1},
		{"a", "b", 0},
		{"ab", "ab", 1},
		{"abb", "ab", 1},
		{"abb", "ab", 1},
		{"aaabaaaabaaabaaaabaaaabaaaabaaaaba", "aaaba", 5},
	}
	for _, tt := range tests {
		t.Run(tt.sequence+"_"+tt.word, func(t *testing.T) {
			if got := maxRepeating(tt.sequence, tt.word); got != tt.want {
				t.Errorf("maxRepeating() = %v, want %v", got, tt.want)
			}
		})
	}
}
