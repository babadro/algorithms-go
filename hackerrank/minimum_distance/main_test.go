package main

import "testing"

func Test_minimumDistances(t *testing.T) {
	tests := []struct {
		name string
		a    []int32
		want int32
	}{
		{"1", []int32{3, 2, 1, 2, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimumDistances(tt.a); got != tt.want {
				t.Errorf("minimumDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}
