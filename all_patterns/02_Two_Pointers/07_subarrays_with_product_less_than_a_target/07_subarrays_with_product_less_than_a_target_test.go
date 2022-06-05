package _7_subarrays_with_product_less_than_a_target

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findSubarrays(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   [][]int
	}{
		{[]int{2, 5, 3, 10}, 30, [][]int{{2}, {5}, {2, 5}, {3}, {5, 3}, {10}}},
		{[]int{8, 2, 6, 5}, 50, [][]int{{8}, {2}, {8, 2}, {6}, {2, 6}, {5}, {6, 5}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if got := findSubarrays(tt.arr, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
