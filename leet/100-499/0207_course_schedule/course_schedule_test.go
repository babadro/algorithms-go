package _207_course_schedule

import "testing"

func Test_canFinish(t *testing.T) {
	tests := []struct {
		name          string
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{"1", 2, [][]int{{1, 0}}, true},
		{"2", 2, [][]int{{1, 0}, {0, 1}}, false},
		{"3", 4, [][]int{{1, 0}, {2, 1}, {3, 2}}, true},
		{"4", 4, [][]int{{1, 0}, {2, 1}, {3, 2}, {2, 3}}, false},
		{"5", 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
		{"6", 1, [][]int{}, true},
		{"7", 3, [][]int{{1, 0}, {0, 2}, {2, 1}}, false},
		{"8", 4, [][]int{{0, 1}, {0, 2}, {1, 3}, {3, 0}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canFinish(tt.numCourses, tt.prerequisites); got != tt.want {
				t.Errorf("canFinish() = %v, want %v", got, tt.want)
			}
		})
	}
}
