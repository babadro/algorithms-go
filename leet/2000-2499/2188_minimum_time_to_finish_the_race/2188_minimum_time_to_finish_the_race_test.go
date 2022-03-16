package _2188_minimum_time_to_finish_the_race

import "testing"

func Test_minimumFinishTime(t *testing.T) {
	tests := []struct {
		tires      [][]int
		changeTime int
		numLaps    int
		want       int
	}{
		{tires: [][]int{{2, 3}, {3, 4}}, changeTime: 5, numLaps: 4, want: 21},
		{tires: [][]int{{1, 10}, {2, 2}, {3, 4}}, changeTime: 6, numLaps: 5, want: 25},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minimumFinishTime(tt.tires, tt.changeTime, tt.numLaps); got != tt.want {
				t.Errorf("minimumFinishTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
