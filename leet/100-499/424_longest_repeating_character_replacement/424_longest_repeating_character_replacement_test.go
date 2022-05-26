package _424_longest_repeating_character_replacement

import (
	"fmt"
	"testing"
)

func Test_characterReplacement(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want int
	}{
		{"ABAB", 2, 4},
		{"AABABBA", 1, 4},
		{"AAABBBB", 0, 4},
		{"ABBCCCDDDD", 0, 4},
		{"AABCDEAA", 0, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s_%d", tt.s, tt.k), func(t *testing.T) {
			if got := characterReplacement(tt.s, tt.k); got != tt.want {
				t.Errorf("characterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
