package _007_reverse_integer

import "testing"

func Test_reverse(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{"1", 123, 321},
		{"2", -123, -321},
		{"3", 120, 21},
		{"4", 12010, 1021},
		{"5", 0, 0},
		{"6", -0, 0},
		{"7", -1, -1},
		{"8", 1, 1},
		{"9", 1534236469, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.x); got != tt.want {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
