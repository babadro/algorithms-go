package _210_course_schedule_2

import (
	"reflect"
	"testing"
)

func TestFindOrder(t *testing.T) {
	cases := []struct {
		numCourses    int
		prerequisites [][]int
	}{
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}},
	}
	_ = cases
	//for i, c := range cases {
	//	res := findOrder(c.numCourses, c.prerequisites)
	//	if
	//}
}

func TestFindOrder2(t *testing.T) {
	t.Log(findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}))
	t.Log(findOrder(3, [][]int{}))
	t.Log(findOrder(2, [][]int{{1, 0}, {0, 1}}))
}

func Test_findOrder3(t *testing.T) {
	tests := []struct {
		vertices int
		edges    [][]int
		want     []int
	}{
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 2, 1, 3}},
		{3, [][]int{{1, 0}, {1, 2}, {0, 1}}, []int{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findOrder3(tt.vertices, tt.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topologicalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
