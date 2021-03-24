package _1801_number_of_orders_in_the_backlog

import (
	"fmt"
	"testing"
)

func Test_getNumberOfBacklogOrders(t *testing.T) {
	tests := []struct {
		orders [][]int
		want   int
	}{
		{[][]int{{10, 5, 0}, {15, 2, 1}, {25, 1, 1}, {30, 4, 0}}, 6},
		{[][]int{{7, 1000000000, 1}, {15, 3, 0}, {5, 999999995, 0}, {5, 1, 1}}, 999999984},
		{[][]int{{19, 28, 0}, {9, 4, 1}, {25, 15, 1}}, 39},
		{[][]int{{1, 29, 1}, {22, 7, 1}, {24, 1, 0}, {25, 15, 1}, {18, 8, 1}, {8, 22, 0}, {25, 15, 1}, {30, 1, 1}, {27, 30, 0}}, 22},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.orders), func(t *testing.T) {
			if got := getNumberOfBacklogOrders(tt.orders); got != tt.want {
				t.Errorf("getNumberOfBacklogOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
