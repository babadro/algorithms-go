package __palindromic_partitioning

import "testing"

func Test_findMPPCuts(t *testing.T) {
	tests := []struct {
		st   string
		want int
	}{
		{"abdbca", 3},
		{"cddpd", 2},
		{"pqr", 2},
		{"pp", 0},
		{"madam", 0},
	}
	for _, tt := range tests {
		t.Run(tt.st, func(t *testing.T) {
			if got := findMPPCutsBottomUp(tt.st); got != tt.want {
				t.Errorf("findMPPCuts() = %v, want %v", got, tt.want)
			}
		})
	}
}
