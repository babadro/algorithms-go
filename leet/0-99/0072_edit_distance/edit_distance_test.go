package _072_edit_distance

import "testing"

func Test_minDistance(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  int
	}{
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"aebgc", "zartbwqc", 5},
		{"", "a", 1},
		{"a", "", 1},
		{"", "", 0},
		{"a", "a", 0},
	}
	for _, tt := range tests {
		t.Run(tt.word1+"_"+tt.word2, func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
