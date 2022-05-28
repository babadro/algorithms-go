package _090_subsets_ii

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/babadro/algorithms-go/slices"
)

func Test_subsetsWithDup(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 2}, [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}},
		{[]int{0}, [][]int{{}, {0}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := subsetsWithDup(tt.nums)
			slices.SortEach(got)
			slices.SortEach(tt.want)
			sort.Slice(got, func(i, j int) bool {
				return slices.Less(got[i], got[j])
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return slices.Less(tt.want[i], tt.want[j])
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsetsWithDup() = %v, want %v", got, tt.want)
			}
		})
	}
}
