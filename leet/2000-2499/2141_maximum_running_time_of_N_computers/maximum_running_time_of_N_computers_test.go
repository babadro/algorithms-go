package _2141_maximum_running_time_of_N_computers

import "testing"

func Test_maxRunTime(t *testing.T) {
	tests := []struct {
		n         int
		batteries []int
		want      int64
	}{
		//{2, []int{3, 3, 3}, 4},
		//{2, []int{1, 1, 1, 1}, 2},
		//{3, []int{10, 10, 3, 5}, 8},
		{6, []int{1, 1, 4, 9, 9, 9}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := maxRunTime(tt.n, tt.batteries); got != tt.want {
				t.Errorf("maxRunTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
