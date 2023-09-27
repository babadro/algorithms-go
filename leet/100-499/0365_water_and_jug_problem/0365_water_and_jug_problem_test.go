package _365_water_and_jug_problem

import "testing"

func Test_canMeasureWater(t *testing.T) {
	tests := []struct {
		jug1Capacity   int
		jug2Capacity   int
		targetCapacity int
		want           bool
	}{
		{3, 5, 4, true},
		{2, 6, 5, false},
		{1, 2, 3, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := canMeasureWater(tt.jug1Capacity, tt.jug2Capacity, tt.targetCapacity); got != tt.want {
				t.Errorf("canMeasureWater() = %v, want %v", got, tt.want)
			}
		})
	}
}
