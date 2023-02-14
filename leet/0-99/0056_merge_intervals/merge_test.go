package _056_merge_intervals

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_merge(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      [][]int
	}{
		{[][]int{{1, 4}, {2, 5}, {7, 9}}, [][]int{{1, 5}, {7, 9}}},
		{[][]int{{1, 4}, {2, 6}, {3, 5}}, [][]int{{1, 6}}},
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}, [][]int{{1, 10}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.intervals), func(t *testing.T) {
			got := merge(tt.intervals)
			sort.Slice(got, func(i, j int) bool {
				return less(got[i], got[j])
			})

			sort.Slice(tt.want, func(i, j int) bool {
				return less(tt.want[i], tt.want[j])
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge() = %v, want %v", got, tt.want)
			}
		})
	}
}

func less(i1, i2 []int) bool {
	if i1[0] != i2[0] {
		return i1[0] < i2[0]
	}

	return i1[1]-i1[0] > i2[1]-i2[0]
}
