package _2332_the_latest_time_to_catch_a_bus

import "testing"

func Test_latestTimeCatchTheBus(t *testing.T) {
	tests := []struct {
		buses      []int
		passengers []int
		capacity   int
		want       int
	}{
		{[]int{10, 20}, []int{2, 17, 18, 19}, 2, 16},
		{[]int{20, 30, 10}, []int{19, 13, 26, 4, 25, 11, 21}, 2, 20},
		{[]int{2}, []int{2}, 2, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := latestTimeCatchTheBus(tt.buses, tt.passengers, tt.capacity); got != tt.want {
				t.Errorf("latestTimeCatchTheBus() = %v, want %v", got, tt.want)
			}
		})
	}
}
