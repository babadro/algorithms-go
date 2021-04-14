package _239_sliding_window_maximum

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_maxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{9, 11}, 2, []int{11}},
		{[]int{4, -2}, 2, []int{4}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %d", tt.nums, tt.k), func(t *testing.T) {
			if got := maxSlidingWindow(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
