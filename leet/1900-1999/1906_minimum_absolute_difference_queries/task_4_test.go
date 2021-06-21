package _1906_minimum_absolute_difference_queries

import (
	"reflect"
	"testing"
)

func Test_minDifference(t *testing.T) {

	tests := []struct {
		nums    []int
		queries [][]int
		want    []int
	}{
		{
			[]int{1, 3, 4, 8},
			[][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}},
			[]int{2, 1, 4, 1},
		},
		{
			[]int{4, 5, 2, 2, 7, 10},
			[][]int{{2, 3}, {0, 2}, {0, 5}, {3, 5}},
			[]int{-1, 1, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minDifference(tt.nums, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
