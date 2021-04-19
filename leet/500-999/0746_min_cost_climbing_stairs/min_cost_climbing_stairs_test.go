package _746_min_cost_climbing_stairs

import "testing"

func Test_minCostClimbingStairs(t *testing.T) {
	tests := []struct {
		name string
		cost []int
		want int
	}{
		{"1", []int{10, 15, 20}, 15},
		{"2", []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{"3", []int{1, 2}, 1},
		{"3", []int{1, 2, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCostClimbingStairs2(tt.cost); got != tt.want {
				t.Errorf("minCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
