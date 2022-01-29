package _2147_numbers_of_ways_to_divide_long_corridor

import "testing"

func Test_numberOfWays(t *testing.T) {
	tests := []struct {
		corridor string
		want     int
	}{
		{"SSPPSPS", 3},
		{"PPSPSP", 1},
		{"S", 0},
		{"SSSS", 1},
		{"SSPSS", 2},
		{"SSPPSS", 3},
		{"SSPPSPPPPPS", 3},
		{"PPPPSSPPSSPPPPP", 3},
		{"P", 0},
		{"SSPSSPSSSPPSPSPPS", 8},
	}
	for _, tt := range tests {
		t.Run(tt.corridor, func(t *testing.T) {
			if got := numberOfWays(tt.corridor); got != tt.want {
				t.Errorf("numberOfWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
