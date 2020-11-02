package main

import "testing"

func Test_bitCount(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"2", 2, 1},
		{"3", 3, 2},
		{"4", 4, 1},
		{"5", 5, 2},
		{"6", 6, 2},
		{"7", 7, 3},
		{"8", 8, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bitCount(tt.num); got != tt.want {
				t.Errorf("bitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
