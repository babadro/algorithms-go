package _207_course_schedule

// doesn't work
func canFinish2(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int]map[int]bool, 0)

	for _, pair := range prerequisites {
		vertex, value := pair[0], pair[1]
		m, ok := graph[vertex]

		if !ok {
			m = make(map[int]bool)
			graph[vertex] = m
		}

		m[value] = true
	}

	if len(graph) == numCourses {
		return false
	}

	for vertex, m := range graph {
		for course := range m {
			if graph[course][vertex] {
				return false
			}
		}
	}

	return true
}
