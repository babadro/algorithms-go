package _227_basic_calculator_II

import "testing"

func Test_calculate(t *testing.T) {
	tests := []struct {
		inputStr string
		want     int
	}{
		//{"3+2*2", 7},
		//{" 3/2 ", 1},
		//{" 3+5 / 2 ", 5},
		{"3+ 2    /10*5 -4/4+ 5* 2", 12},
	}

	for _, tt := range tests {
		t.Run(tt.inputStr, func(t *testing.T) {
			if got := calculate(tt.inputStr); got != tt.want {
				t.Errorf("calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
