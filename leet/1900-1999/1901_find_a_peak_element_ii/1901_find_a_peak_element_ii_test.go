package _901_find_a_peak_element_ii

import (
	"reflect"
	"testing"
)

func Test_findPeakGrid(t *testing.T) {
	tests := []struct {
		mat  [][]int
		want []int
	}{
		{
			[][]int{{1, 4}, {3, 2}}, []int{0, 1},
		},
		{
			[][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}, []int{1, 1},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findPeakGrid(tt.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findPeakGrid() = %v, want %v", got, tt.want)
			}
		})
	}
}
