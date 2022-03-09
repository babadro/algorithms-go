package _2132_stamping_the_grid

import "testing"

func Test_possibleToStamp(t *testing.T) {
	tests := []struct {
		grid        [][]int
		stampHeight int
		stampWidth  int
		want        bool
	}{
		{
			[][]int{
				{1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0}, {1, 0, 0, 0},
			},
			4, 3, true,
		},
		{
			[][]int{
				{1, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 1, 0}, {0, 0, 0, 1},
			},
			2, 2, false,
		},
		{ // doesn't work
			[][]int{
				{1, 1, 1, 0}, {0, 0, 0, 1}, {1, 1, 1, 0}, {1, 1, 0, 0}, {0, 0, 0, 0}, {0, 1, 0, 1}, {0, 1, 0, 1}, {1, 1, 1, 0}, {1, 0, 1, 1}, {0, 0, 0, 1},
			}, 13, 16, false,
		},
		{
			[][]int{
				{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1},
			}, 9, 4, true,
		},
		{
			[][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 0, 1}, {0, 0, 0, 1, 1}},
			2, 2, false,
		},
		{
			[][]int{{1, 0, 0, 1, 1}, {1, 0, 0, 0, 1}, {1, 1, 1, 0, 1}},
			2, 1, true,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := possibleToStamp2(tt.grid, tt.stampHeight, tt.stampWidth); got != tt.want {
				t.Errorf("possibleToStamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
