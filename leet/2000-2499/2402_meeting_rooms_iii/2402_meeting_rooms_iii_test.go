package _2402_meeting_rooms_iii

import (
	"fmt"
	"testing"
)

func Test_mostBooked(t *testing.T) {
	tests := []struct {
		n        int
		meetings [][]int
		want     int
	}{
		{
			2, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}, 0,
		},
		{
			3, [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}}, 1,
		},
		{
			4, [][]int{{18, 19}, {3, 12}, {17, 19}, {2, 13}, {7, 10}}, 0,
		},
		{
			5, [][]int{{40, 47}, {7, 16}, {27, 38}, {16, 43}, {38, 40}, {2, 25}}, 0,
		},
		{100, bigRooms1, 15},
	}
	for _, tt := range tests {
		lenName := len(tt.meetings)
		if lenName > 10 {
			lenName = 10
		}

		t.Run(fmt.Sprintf("%v", tt.meetings[:lenName]), func(t *testing.T) {
			if got := mostBooked(tt.n, tt.meetings); got != tt.want {
				t.Errorf("mostBooked() = %v, want %v", got, tt.want)
			}
		})
	}
}
