package _970_last_day_where_you_can_still_cross

import "testing"

func Test_latestDayToCross(t *testing.T) {

	tests := []struct {
		row   int
		col   int
		cells [][]int
		want  int
	}{
		{2, 2, [][]int{{1, 1}, {2, 1}, {1, 2}, {2, 2}}, 2},
		{2, 2, [][]int{{1, 1}, {1, 2}, {2, 1}, {2, 2}}, 1},
		{3, 3, [][]int{{1, 2}, {2, 1}, {3, 3}, {2, 2}, {1, 1}, {1, 3}, {2, 3}, {3, 2}, {3, 1}}, 3},
		{2, 2, [][]int{{1, 1}, {1, 1}, {1, 1}, {1, 1}}, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := latestDayToCross(tt.row, tt.col, tt.cells); got != tt.want {
				t.Errorf("latestDayToCross() = %v, want %v", got, tt.want)
			}
		})
	}
}
