package _046_permutations

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/babadro/algorithms-go/slices"
)

func Test_permute(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := permute3(tt.nums)
			sort.Slice(got, func(i, j int) bool {
				return slices.Less(got[i], got[j])
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return slices.Less(tt.want[i], tt.want[j])
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("permute2() = %v, want %v", got, tt.want)
			}
		})
	}
}
