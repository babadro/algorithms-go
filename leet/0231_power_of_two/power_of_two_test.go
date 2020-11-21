package _231_power_of_two

import "testing"

func Test_isPowerOfTwo(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"1", 1, true},
		{"2", 2, true},
		{"3", 3, false},
		{"0", 0, false},
		{"4", 4, true},
		{"128", 128, true},
		{"129", 129, false},
		{"-1", -1, false},
		{"-2", -2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.n); got != tt.want {
				t.Errorf("isPowerOfTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
