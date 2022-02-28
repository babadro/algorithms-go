package _2179_count_good_triplets_in_array

import "testing"

func Test_goodTriplets(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int64
	}{
		{[]int{2, 0, 1, 3}, []int{0, 1, 2, 3}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := goodTriplets(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("goodTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
