package _1624_largest_substring_between_two_equal_chars

import "testing"

func Test_maxLengthBetweenEqualCharacters(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"aa", 0},
		{"abca", 2},
		{"cbzxy", -1},
		{"cabbac", 4},
		{"aaaa", 2},
		{"a", -1},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := maxLengthBetweenEqualCharacters(tt.s); got != tt.want {
				t.Errorf("maxLengthBetweenEqualCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
