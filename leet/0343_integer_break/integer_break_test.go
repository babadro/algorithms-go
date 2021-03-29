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
		{4, 4},
		{6, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := integerBreak3(tt.n); got != tt.want {
				t.Errorf("integerBreak() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_2(t *testing.T) {
	for n := 7; n <= 20; n++ {
		t.Log(integerBreak(n))
	}
}
