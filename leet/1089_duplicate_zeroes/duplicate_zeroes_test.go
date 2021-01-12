package _1089_duplicate_zeroes

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_duplicateZeros(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0}, []int{0}},
		{[]int{0, 0, 1}, []int{0, 0, 0}},
		{[]int{1, 0, 0, 2, 3, 4}, []int{1, 0, 0, 0, 0, 2}},
		{[]int{0, 1, 7, 6, 0, 2, 0, 7}, []int{0, 0, 1, 7, 6, 0, 0, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			duplicateZeros2(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Errorf("got = %v, want %v", tt.arr, tt.want)
			}
		})
	}
}
