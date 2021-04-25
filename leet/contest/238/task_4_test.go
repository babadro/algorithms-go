package _38

import (
	"fmt"
	"testing"
)

func Test_maxBuilding(t *testing.T) {
	tests := []struct {
		n            int
		restrictions [][]int
		want         int
	}{
		{5, [][]int{{2, 1}, {4, 1}}, 2},
		{6, [][]int{}, 5},
		{10, [][]int{{5, 3}, {2, 5}, {7, 4}, {10, 3}}, 5},
		{10, [][]int{{8, 5}, {9, 0}, {6, 2}, {4, 0}, {3, 2}, {10, 0}, {5, 3}, {7, 3}, {2, 4}}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.restrictions), func(t *testing.T) {
			if got := maxBuilding(tt.n, tt.restrictions); got != tt.want {
				t.Errorf("maxBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}
