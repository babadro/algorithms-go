package _343_integer_break

import (
	"fmt"
	"testing"
)

func Test_integerBreak(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{2, 1},
		{10, 36},
		{3, 2},
		{6, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := integerBreak(tt.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}
