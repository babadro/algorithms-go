package _031_next_permutation

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_nextPermutation(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1}, []int{1}},
		{[]int{2, 3, 1}, []int{3, 1, 2}},
		{[]int{2, 4, 3, 1}, []int{3, 1, 2, 4}},
		{[]int{5, 4, 7, 5, 3, 2}, []int{5, 5, 2, 3, 4, 7}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			nextPermutation(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("want %v, got %v", tt.want, tt.nums)
			}
		})
	}
}

func Test(t *testing.T) {
	nums := []int{1, 2, 3, 4}

	i := 0
	target := []int{4, 3, 2, 1}
	for {
		if reflect.DeepEqual(nums, target) {
			t.Log(i)

			return
		}

		nextPermutation(nums)
		i++
	}
}
