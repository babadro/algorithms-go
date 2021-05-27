package _866_number_of_ways_to_rearrange_sticks_with_k_sticks_visible

import (
	"fmt"
	"testing"
)

func Test_visibleCount(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2}, 2},
		{[]int{2, 3, 1}, 2},
		{[]int{2, 1, 3}, 2},
		{[]int{1, 2, 3, 4, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := visibleCount(tt.nums); got != tt.want {
				t.Errorf("visibleCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_rearrangeSticks(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{3, 2, 3},
		{5, 5, 1},
		//{20, 11, 647427950},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.n, tt.k), func(t *testing.T) {
			if got := rearrangeSticks(tt.n, tt.k); got != tt.want {
				t.Errorf("rearrangeSticks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bruteforce(t *testing.T) {
	n := 4
	for i := 1; i <= n; i++ {
		t.Logf("n=%d, k=%d: %d", n, i, rearrangeSticks(n, i))
	}
}

func Test_bruteforce1(t *testing.T) {
	rearrangeSticks(4, 3)
}
