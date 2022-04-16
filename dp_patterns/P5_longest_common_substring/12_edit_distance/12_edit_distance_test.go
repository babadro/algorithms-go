package _2_edit_distance

import "testing"

func Test_findMinOperationsBruteForce(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want int
	}{
		{"bat", "but", 1},
		{"abdca", "cbda", 2},
		{"passpot", "ppsspqrt", 3},
	}
	for _, tt := range tests {
		t.Run(tt.s1+"_"+tt.s2, func(t *testing.T) {
			if got := findMinOperationsBottomUp(tt.s1, tt.s2); got != tt.want {
				t.Errorf("findMinOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
