package _0818_race_car

import (
	"strconv"
	"testing"
)

func Test_racecar(t *testing.T) {
	tests := []struct {
		target int
		want   int
	}{
		{3, 2},
		{6, 5},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(tt.target), func(t *testing.T) {
			if got := racecar(tt.target); got != tt.want {
				t.Errorf("racecar() = %v, want %v", got, tt.want)
			}
		})
	}
}
