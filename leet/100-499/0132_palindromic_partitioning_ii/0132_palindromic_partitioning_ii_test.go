package _132_palindromic_partitioning_ii

import "testing"

func Test_minCut(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"aab", 1},
		{"a", 0},
		{"ab", 1},
		{"abdbca", 3},
		{"cddpd", 2},
		{"pqr", 2},
		{"pp", 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minCut(tt.s); got != tt.want {
				t.Errorf("minCut() = %v, want %v", got, tt.want)
			}
		})
	}
}
