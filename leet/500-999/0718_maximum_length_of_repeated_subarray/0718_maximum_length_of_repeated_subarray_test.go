package _718_maximum_length_of_repeated_subarray

import "testing"

func Test_findLength(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findLength(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
