package _2172_maximum_and_sum_of_array

import "testing"

func Test_maximumANDSum(t *testing.T) {
	tests := []struct {
		nums     []int
		numSlots int
		want     int
	}{
		{[]int{1, 2, 3, 4, 5, 6}, 3, 9},
		{[]int{1, 3, 10, 4, 7, 1}, 9, 24},
		{[]int{14, 7, 9, 8, 2, 4, 11, 1, 9}, 8, 40},
		{[]int{10, 5, 3, 6, 11, 8, 8}, 4, 16},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maximumANDSum(tt.nums, tt.numSlots); got != tt.want {
				t.Errorf("maximumANDSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
