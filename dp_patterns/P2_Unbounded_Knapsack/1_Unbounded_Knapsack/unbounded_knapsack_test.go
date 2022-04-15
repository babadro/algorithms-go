package __Unbounded_Knapsack

import (
	"reflect"
	"testing"
)

func Test_knapsack(t *testing.T) {
	tests := []struct {
		profits  []int
		weights  []int
		capacity int
		want     int
	}{
		//{[]int{15, 20, 50}, []int{1, 2, 3}, 5, 80},
		{[]int{10, 9, 8}, []int{1, 1, 1}, 3, 30},
		//{[]int{15, 20, 50, 4, 5, 6, 7, 8}, []int{1, 2, 3, 4, 5, 6, 7, 8}, 130, 2165},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := knapsackBottomUp(tt.profits, tt.weights, tt.capacity); got != tt.want {
				t.Errorf("knapsackBruteForce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindSelectedElements(t *testing.T) {
	tests := []struct {
		dp       [][]int
		weights  []int
		profits  []int
		capacity int
		want     []int
	}{
		{[][]int{
			{0, 10, 20, 30},
			{0, 10, 20, 30},
			{0, 10, 20, 30},
		}, []int{1, 1, 1}, []int{10, 9, 8}, 3, []int{0, 0, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindSelectedElements(tt.dp, tt.weights, tt.profits, tt.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSelectedElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
