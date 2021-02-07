package _27

import (
	"fmt"
	"testing"
)

func Test_maximumScore(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		c    int
		want int
	}{
		{2, 4, 6, 6},
		{4, 4, 6, 7},
		{1, 8, 8, 8},
		{2, 8, 8, 9},
		{0, 2, 2, 2},
		{0, 0, 2, 0},
		{6, 2, 1, 3},
		{1, 1, 1, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d_%d_%d", tt.a, tt.b, tt.c), func(t *testing.T) {
			if got := maximumScore(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("maximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
