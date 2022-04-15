package __0_1_knapsack

import "testing"

func TestKnapsack(t *testing.T) {
	tests := []struct {
		weights  []int
		profits  []int
		capacity int
		want     int
	}{
		{[]int{2, 3, 1, 4}, []int{4, 5, 3, 7}, 5, 10},
		{[]int{1, 2, 3, 5}, []int{1, 6, 10, 16}, 7, 22},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := KnapsackBruteforce(tt.weights, tt.profits, tt.capacity); got != tt.want {
				t.Errorf("KnapsackBruteforce() = %v, want %v", got, tt.want)
			}
			if got := KnapsackTopDown(tt.weights, tt.profits, tt.capacity); got != tt.want {
				t.Errorf("KnapsackTopDown() = %v, want %v", got, tt.want)
			}
			if got := KnapsackBottomUp(tt.weights, tt.profits, tt.capacity); got != tt.want {
				t.Errorf("KnapsackBottomUp() = %v, want %v", got, tt.want)
			}
		})
	}
}
