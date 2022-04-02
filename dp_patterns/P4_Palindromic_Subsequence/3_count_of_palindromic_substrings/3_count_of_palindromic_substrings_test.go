package __count_of_palindromic_substrings

import "testing"

func Test_findCPSBottomUp(t *testing.T) {
	tests := []struct {
		st   string
		want int
	}{
		{"abdbca", 7},
		{"cddpd", 7},
		{"pqr", 3},
	}
	for _, tt := range tests {
		t.Run(tt.st, func(t *testing.T) {
			if got := findCPSBottomUp(tt.st); got != tt.want {
				t.Errorf("findCPSBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
