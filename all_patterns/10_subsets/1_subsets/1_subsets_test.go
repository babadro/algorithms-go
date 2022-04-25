package _1_subsets

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findSubsets(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 3}, [][]int{{}, {1}, {3}, {1, 3}}},
		{[]int{1, 5, 3}, [][]int{{}, {1}, {5}, {3}, {1, 5}, {1, 3}, {5, 3}, {1, 5, 3}}},
		{[]int{1, 2, 3, 4}, [][]int{
			{},
			{1},
			{2},
			{3},
			{4},
			{1, 2},
			{1, 3},
			{1, 4},
			{2, 3},
			{2, 4},
			{3, 4},
			{1, 2, 3},
			{1, 2, 4},
			{1, 3, 4},
			{2, 3, 4},
			{1, 2, 3, 4},
		}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := findSubsetsIterative(tt.nums)

			sort.Slice(got, func(i, j int) bool {
				return less(got[i], got[j])
			})

			sort.Slice(tt.want, func(i, j int) bool {
				return less(tt.want[i], tt.want[j])
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}

func less(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return len(arr1) < len(arr2)
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return arr1[i] < arr2[i]
		}
	}

	return false
}
