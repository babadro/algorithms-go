package _489_robot_cleaner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_cleanRoom(t *testing.T) {
	tests := []struct {
		m              [][]int
		startX, startY int
	}{
		{
			[][]int{{1}}, 0, 0,
		},
		{[][]int{{1, 1, 1, 1}}, 0, 0},
		{
			[][]int{
				{1, 1, 1, 1, 1, 0, 1, 1},
				{1, 1, 1, 1, 1, 0, 1, 1},
				{1, 0, 1, 1, 1, 1, 1, 1},
				{0, 0, 0, 1, 0, 0, 0, 0},
				{1, 1, 1, 1, 1, 1, 1, 1},
			}, 1, 3,
		},
	}
	for _, tt := range tests {
		matrix = tt.m
		realX, realY = tt.startX, tt.startY

		t.Run("", func(t *testing.T) {
			cleanRoom2(&Robot{})

			for y := range matrix {
				for x := range matrix[y] {
					val := matrix[y][x]
					require.True(t, val == 0 || val == 2)
				}
			}
		})
	}
}
