package _001_two_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{0, 4, 3, 0}, 0, []int{0, 3}},
		{[]int{3, 2, 3}, 6, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.nums, tt.target), func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
