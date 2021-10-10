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
		{"abcdabcdabcd", 2, 'd', 2, "dd"},
		{"a", 1, 'a', 1, "a"},
		{"aaaaabddeejfghhhhiilllmmnooqqqrrstuvvwjwxxxyyy", 3, 'j', 2, "ajj"},
		{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", 60, 'a', 60, "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
		{"wuynymkihfdcbabefiiymnoyyytywzy", 16, 'y', 4, "abefiimnoyytywzy"},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := smallestSubsequence(tt.s, tt.k, tt.letter, tt.repetition); got != tt.want {
				t.Errorf("smallestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
