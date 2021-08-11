package _1964_find_the_longest_valid_obstacle_course_at_each_position

// tle, but, probably, correct
func longestObstacleCourseAtEachPosition2(obstacles []int) []int {
	res := make([]int, len(obstacles))

	res[0] = 1
	for i := 1; i < len(res); i++ {
		max := 0
		for j := i - 1; j >= 0 && j+1 > max; j-- {
			if obstacles[j] > obstacles[i] {
				continue
			}

			if res[j] > max {
				max = res[j]
			}

			if res[j] == res[i] {
				break
			}
		}

		res[i] = max + 1
	}

	return res
}
