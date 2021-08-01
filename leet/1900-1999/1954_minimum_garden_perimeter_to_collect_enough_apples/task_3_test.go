package _1954_minimum_garden_perimeter_to_collect_enough_apples

import (
	"strconv"
	"testing"
)

func Test_minimumPerimeter(t *testing.T) {
	tests := []struct {
		neededApples int64
		want         int64
	}{
		{1, 8},
		{12, 8},
		{13, 16},
		{1_000_000_000, 5040},
		{15_000_000_000_000_000, 1242896},
	}
	for _, tt := range tests {
		t.Run(strconv.Itoa(int(tt.neededApples)), func(t *testing.T) {
			if got := minimumPerimeter(tt.neededApples); got != tt.want {
				t.Errorf("minimumPerimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sum(t *testing.T) {
	for i := 1; i <= 5; i++ {
		t.Log(i*i*i - 11*i)
	}
}
