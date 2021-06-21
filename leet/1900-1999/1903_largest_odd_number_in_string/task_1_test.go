package _903_largest_odd_number_in_string

import "testing"

func Test_largestOddNumber(t *testing.T) {
	tests := []struct {
		num  string
		want string
	}{
		{"52", "5"},
		{"4206", ""},
		{"35427", "35427"},
	}
	for _, tt := range tests {
		t.Run(tt.num, func(t *testing.T) {
			if got := largestOddNumber(tt.num); got != tt.want {
				t.Errorf("largestOddNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
