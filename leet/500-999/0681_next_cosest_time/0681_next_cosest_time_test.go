package _681_next_cosest_time

import "testing"

func Test_nextClosestTime(t *testing.T) {
	tests := []struct {
		time string
		want string
	}{
		{"19:34", "19:39"},
		{"23:59", "22:22"},
	}
	for _, tt := range tests {
		t.Run(tt.time, func(t *testing.T) {
			if got := nextClosestTime(tt.time); got != tt.want {
				t.Errorf("nextClosestTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
