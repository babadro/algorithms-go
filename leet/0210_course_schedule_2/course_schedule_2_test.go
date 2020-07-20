package _210_course_schedule_2

import "testing"

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
