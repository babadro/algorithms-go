package __Unbounded_Knapsack

import "testing"

func Test_knapsackBruteForce(t *testing.T) {
	tests := []struct {
		profits  []int
		weights  []int
		capacity int
		want     int
	}{
		{[]int{15, 20, 50}, []int{1, 2, 3}, 5, 80},
		{[]int{15, 20, 50, 4, 5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}, 130, 2165},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := knapsackBottomUp(tt.profits, tt.weights, tt.capacity); got != tt.want {
				t.Errorf("knapsackBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}
