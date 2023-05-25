package _1136_parallel_courses

import "testing"

func Test_minimumSemesters(t *testing.T) {
	tests := []struct {
		n         int
		relations [][]int
		want      int
	}{
		{3, [][]int{{1, 3}, {2, 3}}, 2},
		{3, [][]int{{1, 2}, {2, 3}, {3, 1}}, -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := minimumSemesters(tt.n, tt.relations); got != tt.want {
				t.Errorf("minimumSemesters() = %v, want %v", got, tt.want)
			}
		})
	}
}
