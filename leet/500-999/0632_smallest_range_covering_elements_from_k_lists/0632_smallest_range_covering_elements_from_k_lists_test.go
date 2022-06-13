package _632_smallest_range_covering_elements_from_k_lists

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_smallestRange(t *testing.T) {
	tests := []struct {
		nums [][]int
		want []int
	}{
		{
			[][]int{
				{4, 10, 15, 24, 26},
				{0, 9, 12, 20},
				{5, 18, 22, 30},
			},
			[]int{20, 24},
		},
		{
			[][]int{
				{1, 2, 3},
				{1, 2, 3},
				{1, 2, 3},
			},
			[]int{1, 1},
		},
		{
			[][]int{
				{1, 5, 8}, {4, 12}, {7, 8, 10},
			},
			[]int{4, 7},
		},
		{
			[][]int{{1, 9}, {4, 12}, {7, 10, 16}}, []int{9, 12},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := smallestRange(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("smallestRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
