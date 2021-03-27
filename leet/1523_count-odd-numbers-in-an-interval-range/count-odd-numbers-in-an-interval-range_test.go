package _1523_count_odd_numbers_in_an_interval_range

import (
	"fmt"
	"testing"
)

func Test_countOdds(t *testing.T) {
	tests := []struct {
		low  int
		high int
		want int
	}{
		{3, 7, 3},
		{8, 10, 1},
		{1, 1, 1},
		{2, 2, 0},
		{0, 0, 0},
		{0, 1, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.low, tt.high), func(t *testing.T) {
			if got := countOdds(tt.low, tt.high); got != tt.want {
				t.Errorf("countOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
