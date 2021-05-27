package _869_longer_contiguous_segments_of_ones_than_zeroes

import "testing"

func Test_checkZeroOnes(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"1101", true},
		{"111000", false},
		{"110100010", false},
		{"1", true},
		{"0", false},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := checkZeroOnes(tt.s); got != tt.want {
				t.Errorf("checkZeroOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
