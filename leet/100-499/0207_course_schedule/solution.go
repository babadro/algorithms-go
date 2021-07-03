package _207_course_schedule

// 4ms 5.9
// todo 3 best solution. Need to understand
func canFinish(numCourses int, prerequisites [][]int) bool {
	in := make([]int, numCourses)
	frees := make([][]int, numCourses)
	next := make([]int, 0, numCourses)

	for _, v := range prerequisites {
		in[v[0]]++
		frees[v[1]] = append(frees[v[1]], v[0])
	}

	for i := 0; i < numCourses; i++ {
		if in[i] == 0 {
			next = append(next, i)
		}
	}

	for i := 0; i != len(next); i++ {
		c := next[i]
		v := frees[c]

		for _, w := range v {
			in[w]--

			if in[w] == 0 {
				next = append(next, w)
			}
		}
	}

	return len(next) == numCourses
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
