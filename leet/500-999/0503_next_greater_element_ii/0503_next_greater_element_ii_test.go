package _503_next_greater_element_ii

import (
	"reflect"
	"testing"
)

func Test_nextGreaterElements(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 1}, []int{2, -1, 2}},
		{[]int{1, 2, 3, 4, 3}, []int{2, 3, 4, -1, 4}},
		{[]int{1}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := nextGreaterElements(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nextGreaterElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
