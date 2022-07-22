package _373_find_k_pairs_with_smalles_sums

import (
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/slices"
)

func Test_kSmallestPairs(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		k     int
		want  [][]int
	}{
		{[]int{1, 7, 11}, []int{2, 4, 6}, 3, [][]int{{1, 2}, {1, 4}, {1, 6}}},
		{[]int{1, 1, 2}, []int{1, 2, 3}, 2, [][]int{{1, 1}, {1, 1}}},
		{[]int{1, 2}, []int{3}, 3, [][]int{{1, 3}, {2, 3}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := kSmallestPairs(tt.nums1, tt.nums2, tt.k)
			slices.Sort2DSlice(tt.want)
			slices.Sort2DSlice(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kSmallestPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
