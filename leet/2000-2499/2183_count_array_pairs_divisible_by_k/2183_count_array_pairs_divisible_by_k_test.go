package _2183_count_array_pairs_divisible_by_k

import "testing"

func Test_countPairs(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int64
	}{
		{[]int{1, 2, 3, 4, 5}, 2, 7},
		{[]int{1, 2, 3, 4}, 5, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := countPairs(tt.nums, tt.k); got != tt.want {
				t.Errorf("countPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
