package __fibonacci_numbers

import (
	"fmt"
	"testing"
)

func TestCalculateFibonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{12, 144},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d", tt.n), func(t *testing.T) {
			if got := CalculateFibonacciOptimized(tt.n); got != tt.want {
				t.Errorf("CalculateFibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
