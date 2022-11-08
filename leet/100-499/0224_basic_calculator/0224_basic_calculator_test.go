package _224_basic_calculator

import "testing"

func Test_calculate(t *testing.T) {

	tests := []struct {
		s    string
		want int
	}{
		{"1 + 1", 2},
		{" 2-1 + 2 ", 3},
		{"(1+(4+5+2)-3)+(6+8)", 23},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := calculate(tt.s); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
