package _2842_count_k_subsequences_of_a_string_with_maximum_beauty

import "testing"

func Test_countKSubsequencesWithMaxBeauty(t *testing.T) {
	tests := []struct {
		s         string
		subSeqLen int
		want      int
	}{
		{"bcca", 2, 4},
		{"abbcd", 4, 2},
		{"ci", 1, 2},
		{"dd", 2, 0},
		{"fkp", 2, 3},
		{"kjojr", 3, 6},
		{"chrbbazyupvfdjuqdyb", 3, 36},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := countKSubsequencesWithMaxBeauty(tt.s, tt.subSeqLen); got != tt.want {
				t.Errorf("countKSubsequencesWithMaxBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
