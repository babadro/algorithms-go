package _2167_minimum_time_to_remove_all_cars_containing_illegal_goods

import "testing"

func Test_minimumTime(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"1100101", 5},
		{"0010", 2},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minimumTime(tt.s); got != tt.want {
				t.Errorf("minimumTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
