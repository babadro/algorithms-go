package _2087_minimum_cost_homecoming_of_a_robot_in_a_grid

import "testing"

func Test_minCost(t *testing.T) {
	tests := []struct {
		startPos []int
		homePos  []int
		rowCosts []int
		colCosts []int
		want     int
	}{
		{[]int{1, 0}, []int{2, 3}, []int{5, 4, 3}, []int{8, 2, 6, 7}, 18},
		{[]int{0, 0}, []int{0, 0}, []int{5}, []int{26}, 0},
		{
			startPos: []int{5, 5},
			homePos:  []int{5, 2},
			rowCosts: []int{7, 1, 3, 3, 5, 3, 22, 10, 23},
			colCosts: []int{5, 5, 6, 2, 0, 16},
			want:     8,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minCost(tt.startPos, tt.homePos, tt.rowCosts, tt.colCosts); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
