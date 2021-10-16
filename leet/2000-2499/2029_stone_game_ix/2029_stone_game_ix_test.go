package _2029_stone_game_ix

import "testing"

func Test_stoneGameIX(t *testing.T) {
	tests := []struct {
		stones []int
		want   bool
	}{
		{[]int{2, 1}, true},
		{[]int{2}, false},
		{[]int{5, 1, 2, 4, 3}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := stoneGameIX(tt.stones); got != tt.want {
				t.Errorf("stoneGameIX() = %v, want %v", got, tt.want)
			}
		})
	}
}
