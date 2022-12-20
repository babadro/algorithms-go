package _167_two_sum_ii_input_array_is_sorted

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	tests := []struct {
		numbers []int
		target  int
		want    []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{1, 2}},
		{[]int{2, 3, 4}, 6, []int{1, 3}},
		{[]int{-1, 0}, -1, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.numbers, tt.target), func(t *testing.T) {
			if got := twoSum(tt.numbers, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
