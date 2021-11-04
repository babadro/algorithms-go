package _045_jump_game_ii

import "testing"

func Test_jump(t *testing.T) {
	tests := []struct {
		jumps []int
		want  int
	}{
		{[]int{2, 1, 1, 1, 4}, 3},
		{[]int{1, 1, 3, 6, 9, 3, 0, 1, 3}, 4},
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{3}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := jumpBestSolution(tt.jumps); got != tt.want {
				t.Errorf("jump() = %v, want %v", got, tt.want)
			}
		})
	}
}
