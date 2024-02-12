package _2366_minimum_replacements_to_sort_the_array

import "testing"

func Test_minimumReplacement(t *testing.T) {
	tests := []struct {
		nums []int
		want int64
	}{
		{[]int{3, 9, 3}, 2},
		{[]int{1, 2, 3, 4, 5}, 0},
		{[]int{12, 9, 7, 6, 17, 19, 21}, 6},
		{[]int{12, 10, 7, 6, 17, 19, 21}, 9},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minimumReplacement(tt.nums); got != tt.want {
				t.Errorf("minimumReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
