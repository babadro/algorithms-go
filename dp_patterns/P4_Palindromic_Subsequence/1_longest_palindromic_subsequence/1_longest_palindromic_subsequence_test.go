package __longest_palindromic_subsequence

import "testing"

func Test_findLPSLen(t *testing.T) {
	tests := []struct {
		st   string
		want int
	}{
		{"abdbca", 5},
		{"cddpd", 3},
		{"pqr", 1},
	}
	for _, tt := range tests {
		t.Run(tt.st, func(t *testing.T) {
			if got := FindLPSLenBottomUp(tt.st); got != tt.want {
				t.Errorf("findLPSLenBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
