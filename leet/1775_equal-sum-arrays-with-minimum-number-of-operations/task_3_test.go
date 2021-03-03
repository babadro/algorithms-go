package _775_equal_sum_arrays_with_minimum_number_of_operations

import "testing"

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name  string
		nums1 []int
		nums2 []int
		want  int
	}{
		{
			"1",
			[]int{1, 2, 3, 4, 5, 6},
			[]int{1, 1, 2, 2, 2, 2},
			3,
		},
		{
			"2",
			[]int{1, 1, 1, 1, 1, 1, 1},
			[]int{6},
			-1,
		},
		{
			"3",
			[]int{6, 6},
			[]int{1},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minOperations(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
