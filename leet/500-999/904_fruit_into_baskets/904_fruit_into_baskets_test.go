package _904_fruit_into_baskets

import (
	"fmt"
	"testing"
)

func Test_totalFruit(t *testing.T) {
	tests := []struct {
		fruits []int
		want   int
	}{
		{[]int{1, 2, 1}, 3},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 3, 2, 2}, 4},
		{[]int{}, 0},
		{[]int{0, 1, 1}, 3},
		{[]int{1, 0, 1, 4, 1, 4, 1, 2, 3}, 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.fruits), func(t *testing.T) {
			if got := totalFruit(tt.fruits); got != tt.want {
				t.Errorf("totalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
