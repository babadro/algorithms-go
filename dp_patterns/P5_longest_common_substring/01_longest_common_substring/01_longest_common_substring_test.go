package _1_longest_common_substring

import (
	"fmt"
	"testing"
)

func Test_findLCSLen(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want int
	}{
		{"abdca", "cbda", 2},
		{"passport", "ppsspt", 3},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.s1, tt.s2), func(t *testing.T) {
			if got := findLCSLenTopDown(tt.s1, tt.s2); got != tt.want {
				t.Errorf("findLCSLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
