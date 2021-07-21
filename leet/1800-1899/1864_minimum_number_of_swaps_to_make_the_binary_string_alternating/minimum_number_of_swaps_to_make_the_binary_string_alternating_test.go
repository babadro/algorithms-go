package _1864_minimum_number_of_swaps_to_make_the_binary_string_alternating

import "testing"

func Test_minSwaps(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"111000", 1},
		{"010", 0},
		{"1110", -1},
		{"11", -1},
		{"00", -1},
		{"1", 0},
		{"0", 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minSwaps(tt.s); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
