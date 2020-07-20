package _210_course_schedule_2

const (
	white byte = iota
	gray
	black
)

// 98% 43%
func findOrder(numCourses int, prerequisites [][]int) []int {
	isPossible := true
	color := make([]byte, numCourses)
	adjList := make([][]int, numCourses)
	topologicalOrder := make([]int, 0, numCourses)

	for i := 0; i < len(prerequisites); i++ {
		dest := prerequisites[i][0]
		src := prerequisites[i][1]
		adjList[src] = append(adjList[src], dest)
	}

	for i := 0; i < numCourses; i++ {
		if color[i] == white {
			dfs(i, adjList, &isPossible, &topologicalOrder, color)
		}
	}

	if !isPossible {
		return nil
	}

	order := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		order[i] = topologicalOrder[numCourses-i-1]
	}

	return order
}

func dfs(node int, adjList [][]int, isPossible *bool, topologicalOrder *[]int, color []byte) {
	if !*isPossible {
		return
	}

	color[node] = gray
	for _, neighbor := range adjList[node] {
		if color[neighbor] == white {
			dfs(neighbor, adjList, isPossible, topologicalOrder, color)
		} else if color[neighbor] == gray {
			*isPossible = false
		}
	}

	color[node] = black
	*topologicalOrder = append(*topologicalOrder, node)
}
