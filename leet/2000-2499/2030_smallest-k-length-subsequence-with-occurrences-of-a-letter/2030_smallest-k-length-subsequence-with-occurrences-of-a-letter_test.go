package _2030_smallest_k_length_subsequence_with_occurrences_of_a_letter

import "testing"

func Test_smallestSubsequenceBruteForce(t *testing.T) {
	tests := []struct {
		s          string
		k          int
		letter     byte
		repetition int
		want       string
	}{
		{s: "leet", k: 3, letter: 'e', repetition: 1, want: "eet"},
		{"leetcode", 4, 'e', 2, "ecde"},
		{"bb", 2, 'b', 2, "bb"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := smallestSubsequenceBruteForce(tt.s, tt.k, tt.letter, tt.repetition); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
