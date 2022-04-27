package _3_squaring_sorted_array

import (
	"reflect"
	"testing"
)

func Test_makeSquares(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{-2, -1, 0, 2, 3}, []int{0, 1, 4, 4, 9}},
		{[]int{1, 2, 3, 4}, []int{1, 4, 9, 16}},
		{[]int{0, 1, 2}, []int{0, 1, 4}},
		{[]int{-1, 1}, []int{1, 1}},
		{[]int{-3, -2, -1, 0, 1, 2, 3}, []int{0, 1, 1, 4, 4, 9, 9}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := makeSquares(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
