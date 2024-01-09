package _253_meeting_rooms_ii

import "testing"

func Test_minMeetingRooms(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{[][]int{{7, 10}, {2, 4}}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minMeetingRooms(tt.intervals); got != tt.want {
				t.Errorf("minMeetingRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
