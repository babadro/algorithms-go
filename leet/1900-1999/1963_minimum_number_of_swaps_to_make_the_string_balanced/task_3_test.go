package _963_minimum_number_of_swaps_to_make_the_string_balanced

import "testing"

func Test_minSwaps(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"][][", 1},
		{"]]][[[", 2},
		{"[]", 0},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := minSwaps(tt.s); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
