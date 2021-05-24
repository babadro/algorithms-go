package _870_minimum_speed_to_arrive_on_time

import (
	"fmt"
	"testing"
)

func Test_minSpeedOnTime(t *testing.T) {
	tests := []struct {
		dist []int
		hour float64
		want int
	}{
		{[]int{1, 3, 2}, 6, 1},
		{[]int{1, 3, 2}, 2.7, 3},
		{[]int{1, 3, 2}, 1.9, -1},
		{[]int{1, 1, 100_000}, 2.01, 10_000_000},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("dist: %v, hour: %f", tt.dist, tt.hour), func(t *testing.T) {
			if got := minSpeedOnTime(tt.dist, tt.hour); got != tt.want {
				t.Errorf("minSpeedOnTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
