package _342_power_of_four

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"1", 1, true},
		{"4", 4, true},
		{"16", 16, true},
		{"64", 64, true},
		{"256", 256, true},
		{"128", 128, false},
		{"8", 128, false},
		{"5", 5, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour2(tt.num); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
