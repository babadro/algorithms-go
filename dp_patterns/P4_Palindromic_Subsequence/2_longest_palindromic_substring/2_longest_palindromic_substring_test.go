package __longest_palindromic_substring

import "testing"

func Test_findLPSLength(t *testing.T) {
	tests := []struct {
		st   string
		want int
	}{
		{"abdbca", 3},
		{"cddpd", 3},
		{"pqr", 1},
	}
	for _, tt := range tests {
		t.Run(tt.st, func(t *testing.T) {
			if got := findLPSLenTopDown(tt.st); got != tt.want {
				t.Errorf("findLPSLenBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
