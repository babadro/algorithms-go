package _1851_minimum_interval_to_include_each_query

import (
	"reflect"
	"testing"
)

func Test_minInterval(t *testing.T) {
	tests := []struct {
		intervals [][]int
		queries   []int
		want      []int
	}{
		{
			[][]int{{1, 4}, {2, 4}, {3, 6}, {4, 4}},
			[]int{2, 3, 4, 5},
			[]int{3, 3, 1, 4},
		},
		{
			[][]int{{2, 3}, {2, 5}, {1, 8}, {20, 25}},
			[]int{2, 19, 5, 22},
			[]int{2, -1, 4, 6},
		},
		{
			[][]int{{4, 5}, {5, 8}, {1, 9}, {8, 10}, {1, 6}},
			[]int{7, 9, 3, 9, 3},
			[]int{4, 3, 6, 3, 6},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minInterval(tt.intervals, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
