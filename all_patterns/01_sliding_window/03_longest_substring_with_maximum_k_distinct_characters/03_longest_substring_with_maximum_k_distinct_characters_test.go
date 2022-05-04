package _3_longest_substring_with_maximum_k_distinct_characters

import (
	"fmt"
	"testing"
)

func Test_findLength(t *testing.T) {
	tests := []struct {
		str  string
		k    int
		want int
	}{
		{"araaci", 2, 4},
		{"araaci", 1, 2},
		{"cbbebi", 3, 5},
		{"cbbebi", 10, 6},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %d", tt.str, tt.k), func(t *testing.T) {
			if got := findLength(tt.str, tt.k); got != tt.want {
				t.Errorf("findLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
