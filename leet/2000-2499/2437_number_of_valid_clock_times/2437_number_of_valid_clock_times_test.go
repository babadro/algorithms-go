package _2437_number_of_valid_clock_times

import "testing"

func Test_countTime(t *testing.T) {
	tests := []struct {
		time string
		want int
	}{
		{"?5:00", 2},
		{"0?:0?", 100},
		{"??:??", 1440},
	}
	for _, tt := range tests {
		t.Run(tt.time, func(t *testing.T) {
			if got := countTime(tt.time); got != tt.want {
				t.Errorf("countTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
