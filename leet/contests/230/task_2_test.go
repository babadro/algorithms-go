package _30

import (
	"fmt"
	"testing"
)

func Test_closestCost(t *testing.T) {
	tests := []struct {
		baseCosts    []int
		toppingCosts []int
		target       int
		want         int
	}{
		{[]int{1, 7}, []int{3, 4}, 10, 10},
		{[]int{2, 3}, []int{4, 5, 100}, 18, 17},
		{[]int{3, 10}, []int{2, 5}, 9, 8},
		{[]int{10}, []int{1}, 1, 10},
		{[]int{9015, 4152, 7816, 5153, 1641, 3402, 5201}, []int{650, 447, 173, 4843}, 9775, 9788},
		{[]int{20, 20}, []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 10}, 30, 30}, // todo 2 how does it work?
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v %v %d", tt.baseCosts, tt.toppingCosts, tt.target)
		t.Run(name, func(t *testing.T) {
			if got := closestCost2(tt.baseCosts, tt.toppingCosts, tt.target); got != tt.want {
				t.Errorf("closestCost() = %v, want %v", got, tt.want)
			}
		})
	}
}

/*
[9015,4152,7816,5153,1641,3402,5201]
[650,447,173,4843]
9775

want 9788, got 9808

4152+4843+447+173+173 = 9788

9015+447+173+173 = 9808
*/
