package _1850_minimum_adjacent_swaps_to_reach_the_kth_smallest_number

import (
	"fmt"
	"testing"
)

func Test_getMinSwaps(t *testing.T) {
	tests := []struct {
		num  string
		k    int
		want int
	}{
		{"5489355142", 4, 2},
		{"11112", 4, 4},
		{"00123", 1, 1},
		{"1234", 23, 6},
		{"5555555555555555555555555555555555555555555555555555555555555555555555555559", 75, 75},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, %d", tt.num, tt.k), func(t *testing.T) {
			if got := getMinSwaps(tt.num, tt.k); got != tt.want {
				t.Errorf("getMinSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
"5555555555555555555555555555555555555555555555555555555555555555555555555559"
75

*/
