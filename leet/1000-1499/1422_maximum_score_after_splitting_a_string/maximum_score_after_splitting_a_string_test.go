package _1422_maximum_score_after_splitting_a_string

import "testing"

func Test_maxScore(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"011101", 5},
		{"00111", 5},
		{"1111", 3},
		{"00", 1},
		{"01", 2},
		{"11", 1},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := maxScore2(tt.s); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
