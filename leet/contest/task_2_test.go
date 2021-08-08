package contest

import (
	"fmt"
	"testing"
)

func Test_minStoneSum(t *testing.T) {
	tests := []struct {
		piles []int
		k     int
		want  int
	}{
		{[]int{5, 4, 9}, 2, 12},
		{[]int{4, 3, 6, 7}, 3, 12},
		{[]int{1}, 100, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v", tt.piles, tt.k), func(t *testing.T) {
			if got := minStoneSum(tt.piles, tt.k); got != tt.want {
				t.Errorf("minStoneSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
