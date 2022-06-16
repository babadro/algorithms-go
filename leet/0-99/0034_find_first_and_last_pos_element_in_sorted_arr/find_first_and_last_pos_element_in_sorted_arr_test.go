package _034_find_first_and_last_pos_element_in_sorted_arr

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_searchRange2(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
		{[]int{1}, 1, []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := searchRange2(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange2() = %v, want %v", got, tt.want)
			}
		})
	}
}
