package _0435_non_overlapping_intervals

import (
	"fmt"
	"testing"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 2}, {2, 3}}, 0},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {-100, -2}, {5, 7}}, 0},
		{tleInput1, 187},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.intervals), func(t *testing.T) {
			if got := eraseOverlapIntervals(tt.intervals); got != tt.want {
				t.Errorf("eraseOverlapIntervals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test1(t *testing.T) {

}
