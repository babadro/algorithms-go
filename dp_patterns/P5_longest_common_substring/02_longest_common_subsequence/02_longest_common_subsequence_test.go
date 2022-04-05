package _2_longest_common_subsequence

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
		{"abdca", "cbda", 3},
		{"passport", "ppsspt", 5},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, %s", tt.s1, tt.s2), func(t *testing.T) {
			if got := findLCSLenBottomUp(tt.s1, tt.s2); got != tt.want {
				t.Errorf("findLCSLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
