package _210_course_schedule_2

import (
	"reflect"
	"testing"
)

func Test_findOrder(t *testing.T) {
	tests := []struct {
		vertices int
		edges    [][]int
		want     []int
	}{
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 2, 1, 3}},
		{3, [][]int{{1, 0}, {1, 2}, {0, 1}}, nil},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findOrder2(tt.vertices, tt.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topologicalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
