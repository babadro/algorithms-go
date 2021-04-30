package utils

import (
	"fmt"
	"testing"
)

func TestLog(t *testing.T) {
	tests := []struct {
		b    int
		x    int
		want int
	}{
		{10, 3050, 3},
		{10, 0, 0},
		{10, 3, 0},
		{10, 9, 0},
		{10, 10, 1},
		{10, 11, 1},
		{2, 8, 3},
		{3, 10, 2},
		{10, 3413231233421, 12},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d, %d", tt.b, tt.x), func(t *testing.T) {
			if got := Log(tt.b, tt.x); got != tt.want {
				t.Errorf("log() = %v, want %v", got, tt.want)
			}
		})
	}
}
