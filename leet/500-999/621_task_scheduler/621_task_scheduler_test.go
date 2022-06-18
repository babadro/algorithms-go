package _621_task_scheduler

import "testing"

func Test_leastInterval(t *testing.T) {
	tests := []struct {
		tasks []byte
		n     int
		want  int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 0, 6},
		{[]byte{'A', 'A', 'A', 'A', 'A', 'A', 'B', 'C', 'D', 'E', 'F', 'G'}, 2, 16},
	}
	for _, tt := range tests {
		t.Run(string(tt.tasks), func(t *testing.T) {
			if got := leastInterval(tt.tasks, tt.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
