package _207_course_schedule

// dyx. passed. best solution. easy to understand
func canFinish(numCourses int, prerequisites [][]int) bool {
	g := make([][]int, numCourses)
	inDegrees := make([]int, numCourses)

	for _, p := range prerequisites {
		src, dst := p[1], p[0]
		g[src] = append(g[src], dst)
		inDegrees[dst]++
	}

	sources := make([]int, 0, numCourses)
	for i, degree := range inDegrees {
		if degree == 0 {
			sources = append(sources, i)
		}
	}

	for i := 0; i < len(sources); i++ {
		src := sources[i]
		for _, vertex := range g[src] {
			inDegrees[vertex]--
			if inDegrees[vertex] == 0 {
				sources = append(sources, vertex)
			}
		}
	}

	return len(sources) == numCourses
}

// https://leetcode.com/problems/course-schedule/discuss/331722/go-(golang)%3A-DFS-and-BFS
// 33% 7% but relatively easy to understand
// todo 2 need to remember
func canFinish3(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int, numCourses)
	for _, p := range prerequisites {
		graph[p[1]] = append(graph[p[1]], p[0])
	}

	state := make([]int, numCourses)

	for i := 0; i < numCourses; i++ {
		if !DFS(graph, i, state) {
			return false
		}
	}

	return true
}

func DFS(g map[int][]int, cur int, state []int) bool {
	if state[cur] == 1 {
		return true
	}

	if state[cur] == 2 {
		return false
	}

	state[cur] = 2

	for _, neighbor := range g[cur] {
		if !DFS(g, neighbor, state) { // have circle
			return false
		}
	}

	state[cur] = 1

	return true
}
