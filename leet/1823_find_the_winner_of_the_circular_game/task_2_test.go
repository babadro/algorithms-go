package _1823_find_the_winner_of_the_circular_game

import (
	"fmt"
	"testing"
)

func Test_findTheWinner(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{5, 2, 3},
		{1, 1, 1},
		{2, 1, 2},
		{3, 3, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d %d", tt.n, tt.k), func(t *testing.T) {
			if got := findTheWinner(tt.n, tt.k); got != tt.want {
				t.Errorf("findTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
