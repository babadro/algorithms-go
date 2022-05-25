package _018_4sum

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/slices"
)

func Test_fourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
		{[]int{2, 0, -1, 1, -2, 2}, 2, [][]int{{-2, 0, 2, 2}, {-1, 0, 1, 2}}},
		{[]int{4, 1, 2, -1, 1, -3}, 1, [][]int{{-3, -1, 1, 4}, {-3, 1, 1, 2}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := fourSum(tt.nums, tt.target)
			slices.SortEach(tt.want)
			slices.SortEach(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
