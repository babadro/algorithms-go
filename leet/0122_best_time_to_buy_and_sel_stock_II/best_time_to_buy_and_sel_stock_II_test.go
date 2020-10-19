package _122_best_time_to_buy_and_sel_stock_II

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{"1", []int{7, 1, 5, 3, 6, 4}, 7},
		{"2", []int{1, 2, 3, 4, 5}, 4},
		{"3", []int{7, 6, 4, 3, 1}, 0},
		{"4", []int{7, 1, 1, 1, 5, 3, 6, 4, 4, 4}, 7},
		{"5", []int{1, 1, 1}, 0},
		{"6", []int{1, 2, 2}, 1},
		{"7", []int{1, 2, 3}, 2},
		{"8", []int{3, 2, 2}, 0},
		{"9", []int{3, 2, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit2(tt.prices); got != tt.want {
				t.Errorf("maxProfit2() = %v, want %v", got, tt.want)
			}
		})
	}
}
