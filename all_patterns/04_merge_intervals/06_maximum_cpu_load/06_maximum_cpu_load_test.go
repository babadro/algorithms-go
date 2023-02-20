package _6_maximum_cpu_load

import "testing"

func Test_findMaxCPULoad(t *testing.T) {
	tests := []struct {
		jobs [][3]int
		want int
	}{
		{[][3]int{{6, 7, 10}, {2, 4, 11}, {8, 12, 15}}, 15},
		{[][3]int{{1, 4, 2}, {2, 4, 1}, {3, 6, 5}}, 8},
		{[][3]int{{1, 4, 3}, {2, 5, 4}, {7, 9, 6}}, 7},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findMaxCPULoad2(tt.jobs); got != tt.want {
				t.Errorf("findMaxCPULoad() = %v, want %v", got, tt.want)
			}
		})
	}
}
