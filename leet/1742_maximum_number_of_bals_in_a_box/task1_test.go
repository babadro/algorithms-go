package _1742_maximum_number_of_bals_in_a_box

import (
	"fmt"
	"testing"
)

func Test_digitsSum(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{1, 1},
		{10, 1},
		{101, 2},
		{499, 22},
		{0, 0},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.num), func(t *testing.T) {
			if got := digitsSum(tt.num); got != tt.want {
				t.Errorf("digitsSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countBalls(t *testing.T) {
	tests := []struct {
		lowLimit  int
		highLimit int
		want      int
	}{
		{1, 10, 2},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %v", tt.lowLimit, tt.highLimit), func(t *testing.T) {
			if got := countBalls(tt.lowLimit, tt.highLimit); got != tt.want {
				t.Errorf("countBalls() = %v, want %v", got, tt.want)
			}
		})
	}
}
