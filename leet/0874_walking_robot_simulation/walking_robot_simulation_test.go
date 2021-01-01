package _874_walking_robot_simulation

import "testing"

func Test_robotSim(t *testing.T) {
	tests := []struct {
		name      string
		commands  []int
		obstacles [][]int
		want      int
	}{
		{
			name:      "1",
			commands:  []int{4, -1, 3},
			obstacles: [][]int{},
			want:      25,
		},
		{
			name:      "2",
			commands:  []int{4, -1, 4, -2, 4},
			obstacles: [][]int{{2, 4}},
			want:      65,
		},
		{
			name:      "3",
			commands:  []int{8},
			obstacles: [][]int{},
			want:      64,
		},
		{
			name:      "4",
			commands:  []int{-1, 8},
			obstacles: [][]int{},
			want:      64,
		},
		{
			name:      "4",
			commands:  []int{-1, 8},
			obstacles: [][]int{{1, 0}},
			want:      0,
		},
		{
			name:      "5",
			commands:  []int{-2, 8, 3, 7, -1},
			obstacles: [][]int{{-4, -1}, {1, -1}, {1, 4}, {5, 0}, {4, 5}, {-2, -1}, {2, -5}, {5, 1}, {-3, -1}, {5, -3}},
			want:      324,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := robotSim2(tt.commands, tt.obstacles); got != tt.want {
				t.Errorf("robotSim() = %v, want %v", got, tt.want)
			}
		})
	}
}
