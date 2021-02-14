package _1759_count_number_of_homogenous_substrings

import "testing"

func Test_countHomogenous(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"abbcccaa", 13},
		{"xy", 2},
		{"zzzzz", 15},
		{"a", 1},
		{"abcd", 4},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := countHomogenous(tt.s); got != tt.want {
				t.Errorf("countHomogenous() = %v, want %v", got, tt.want)
			}
		})
	}
}
