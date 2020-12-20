package _637_widest_vertical_area_between_two_points_containing_no_points

import "testing"

func Test_maxWidthOfVerticalArea(t *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   int
	}{
		{
			name:   "1",
			points: [][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}},
			want:   1,
		},
		{
			name:   "2",
			points: [][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}},
			want:   3,
		},
		{
			name:   "3",
			points: [][]int{{0, 1}, {0, 2}},
			want:   0,
		},
		{
			name:   "4",
			points: [][]int{{-2, -2}, {-1, 2}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxWidthOfVerticalArea(tt.points); got != tt.want {
				t.Errorf("maxWidthOfVerticalArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
