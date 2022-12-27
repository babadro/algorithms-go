package _454_4sum_ii

import "testing"

func Test_fourSumCount(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		nums3 []int
		nums4 []int

		want int
	}{
		{
			[]int{1, 2}, []int{-2 - 1}, []int{-1, 2}, []int{0, 2},
			2,
		},
		{
			[]int{0}, []int{0}, []int{0}, []int{0}, 1,
		},
		{
			[]int{0, 1, -1},
			[]int{-1, 1, 0},
			[]int{0, 0, 1},
			[]int{-1, 1, 1},
			17,
		},
		{nums1, nums2, nums3, nums4, 4294623},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := fourSumCount(tt.nums1, tt.nums2, tt.nums3, tt.nums4); got != tt.want {
				t.Errorf("fourSumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
