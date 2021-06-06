package _1889_minimum_space_wasted_from_packaging

import "testing"

func Test_minWastedSpace(t *testing.T) {
	tests := []struct {
		packages []int
		boxes    [][]int
		want     int
	}{
		{[]int{2, 3, 5}, [][]int{{4, 8}, {2, 8}}, 6},
		{[]int{2, 3, 5}, [][]int{{1, 4}, {2, 3}, {3, 4}}, -1},
		{[]int{3, 5, 8, 10, 11, 12}, [][]int{{12}, {11, 9}, {10, 5, 14}}, 9},
		{tle1Packages, tle1Boxes, 499949986},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minWastedSpace(tt.packages, tt.boxes); got != tt.want {
				t.Errorf("minWastedSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
