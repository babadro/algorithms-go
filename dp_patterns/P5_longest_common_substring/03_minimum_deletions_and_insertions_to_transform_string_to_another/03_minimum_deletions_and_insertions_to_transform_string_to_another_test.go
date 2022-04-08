package _3_minimum_deletions_and_insertions_to_transform_string_to_another

import (
	"fmt"
	"testing"
)

func Test_findMDI(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want int
	}{
		{"abc", "fbc", 2},
		{"abdca", "cbda", 3},
		{"passport", "ppsspt", 4},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%s, %s", tt.s1, tt.s2), func(t *testing.T) {
			if got := findMDI(tt.s1, tt.s2); got != tt.want {
				t.Errorf("findMDI() = %v, want %v", got, tt.want)
			}
		})
	}
}
