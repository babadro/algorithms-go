package _1_top_k_numbers

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findKLargestNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{3, 1, 5, 12, 2, 11}, 3, []int{5, 12, 11}},
		{[]int{5, 12, 11, -1, 12}, 3, []int{12, 11, 12}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := findKLargestNumbers(tt.nums, tt.k)
			sort.Ints(got)
			sort.Ints(tt.want)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findKLargestNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
