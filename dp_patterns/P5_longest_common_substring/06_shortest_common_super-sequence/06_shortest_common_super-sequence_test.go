package _6_shortest_common_super_sequence

import (
	"fmt"
	"testing"
)

func Test_findSCSLen(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want int
	}{
		{"abcf", "bcdf", 5},
		{"dynamic", "programming", 15},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s %s", tt.s1, tt.s2), func(t *testing.T) {
			if got := findSCSLenBottomUp(tt.s1, tt.s2); got != tt.want {
				t.Errorf("findSCSLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
