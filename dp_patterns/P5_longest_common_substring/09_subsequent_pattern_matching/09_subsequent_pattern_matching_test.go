package _9_subsequent_pattern_matching

import (
	"testing"
)

func Test_findSPMCount(t *testing.T) {
	tests := []struct {
		st   string
		pat  string
		want int
	}{
		{"baxmx", "ax", 2},
		{"tomorrow", "tor", 4},
	}
	for _, tt := range tests {
		t.Run(tt.st+"_"+tt.pat, func(t *testing.T) {
			if got := findSPMCountBottomUp(tt.st, tt.pat); got != tt.want {
				t.Errorf("findSPMCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
