package _1833_maximum_ice_cream_bars

import (
	"fmt"
	"testing"
)

func Test_maxIceCream(t *testing.T) {
	tests := []struct {
		costs []int
		coins int
		want  int
	}{
		{[]int{1, 3, 2, 4, 1}, 7, 4},
		{[]int{10, 6, 8, 7, 7, 8}, 5, 0},
		{[]int{1, 6, 3, 1, 2, 5}, 20, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.costs), func(t *testing.T) {
			if got := maxIceCream(tt.costs, tt.coins); got != tt.want {
				t.Errorf("maxIceCream() = %v, want %v", got, tt.want)
			}
		})
	}
}
