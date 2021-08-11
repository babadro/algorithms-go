package _1964_find_the_longest_valid_obstacle_course_at_each_position

import "sort"

func longestObstacleCourseAtEachPosition(obstacles []int) []int {
	var seq []int
	res := make([]int, len(obstacles))
	for i, ob := range obstacles {
		n := len(seq)
		idx := sort.Search(n, func(i int) bool {
			return seq[i] > ob
		})

		if idx == n {
			seq = append(seq, ob)
		} else {
			seq[idx] = ob
		}

		res[i] = idx + 1
	}

	return res
}
