package _004_median_of_two_sorted_arrays

import (
	"fmt"
	"testing"
)

func Test_findMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{[]int{1, 4}, []int{2, 3}, 2.5},
		{[]int{0, 0}, []int{0, 0}, 0},
		{[]int{}, []int{1}, 1},
		{[]int{2}, []int{}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", tt.nums1, tt.nums2), func(t *testing.T) {
			if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
