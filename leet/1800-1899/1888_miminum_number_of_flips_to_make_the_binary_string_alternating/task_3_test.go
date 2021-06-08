package _1888_miminum_number_of_flips_to_make_the_binary_string_alternating

import "testing"

func Test_minFlips(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"111000", 2},
		{"010", 0},
		{"1110", 1},
		{"0", 0},
		{"1", 0},
		{"01001001101", 2},
		{tleInput1, 16358},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minFlips(tt.s); got != tt.want {
				t.Errorf("minFlips() = %v, want %v", got, tt.want)
			}
		})
	}
}
