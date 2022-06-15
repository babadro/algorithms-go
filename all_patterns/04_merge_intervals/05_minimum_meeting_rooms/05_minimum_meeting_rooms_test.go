package _5_minimum_meeting_rooms

import (
	"fmt"
	"testing"
)

func Test_findMinimumMeetingRooms(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{1, 4}, {2, 5}, {7, 9}}, 2},
		{[][]int{{6, 7}, {2, 4}, {8, 12}}, 1},
		{[][]int{{1, 4}, {2, 3}, {3, 6}}, 2},
		{[][]int{{4, 5}, {2, 3}, {2, 4}, {3, 5}}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.intervals), func(t *testing.T) {
			if got := findMinimumMeetingRooms(tt.intervals); got != tt.want {
				t.Errorf("findMinimumMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
