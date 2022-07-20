package _096_unique_binary_search_trees

import (
	"strconv"
	"testing"
)

func Test_numTrees(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{3, 5},
		{1, 1},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.n), func(t *testing.T) {
			if got := numTrees(tt.n); got != tt.want {
				t.Errorf("numTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
