package contest2

import (
	"reflect"
	"sort"
	"testing"
)

func Test_getSums(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, -3}, []int{-3, -2, -1, 0, 0, 1, 2, 3}},
		{[]int{1}, []int{0, 1}},
		{[]int{0, 1}, []int{0, 0, 1, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := getSums(tt.nums)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recoverArray(t *testing.T) {
	tests := []struct {
		n    int
		sums []int
		want []int
	}{
		{1, []int{0, 1}, []int{1}},
		{3, []int{-3, -2, -1, 0, 0, 1, 2, 3}, []int{-3, 1, 2}},
		{2, []int{0, 0, 0, 0}, []int{0, 0}},
		{4, []int{0, 0, 5, 5, 4, -1, 4, 9, 9, -1, 4, 3, 4, 8, 3, 8}, []int{0, -1, 4, 5}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := recoverArray(tt.n, tt.sums)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
