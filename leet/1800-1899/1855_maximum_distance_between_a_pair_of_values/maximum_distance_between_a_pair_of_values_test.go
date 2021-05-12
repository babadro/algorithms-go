package _855_maximum_distance_between_a_pair_of_values

import "testing"

func Test_maxDistance(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{55, 30, 5, 4, 2}, []int{100, 20, 10, 10, 5}, 2},
		{[]int{2, 2, 2}, []int{10, 10, 1}, 1},
		{[]int{30, 29, 19, 5}, []int{25, 25, 25, 25, 25}, 2},
		{[]int{5, 4}, []int{3, 2}, 0},
		{[]int{3, 2, 1}, []int{5, 5, 5}, 2},
		{[]int{1}, []int{2}, 0},
		{[]int{5, 4, 3}, []int{2, 1}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxDistance(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("maxDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
