package _480_sliding_window_median

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_medianSlidingWindow(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []float64
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []float64{1, -1, -1, 3, 5, 6}},
		{[]int{1, 2, 3, 4, 2, 3, 1, 4, 2}, 3, []float64{2, 3, 3, 3, 2, 3, 2}},
		{[]int{1}, 1, []float64{1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := medianSlidingWindow(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("medianSlidingWindow() = %v, want %v", got, tt.want)
			}
		})
	}
}
