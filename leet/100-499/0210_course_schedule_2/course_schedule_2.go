package _210_course_schedule_2

// tptl. passed. simple
func findOrder3(vertices int, edges [][]int) []int {
	inDegree := make([]int, vertices)
	for _, edge := range edges {
		inDegree[edge[0]]++
	}

	// build the graph
	g := make([][]int, vertices)
	for _, edge := range edges {
		src, dst := edge[1], edge[0]
		g[src] = append(g[src], dst)
	}

	var sourcesQueue []int
	for vertex := range g {
		if inDegree[vertex] == 0 {
			sourcesQueue = append(sourcesQueue, vertex)
		}
	}

	var res []int
	for len(sourcesQueue) > 0 {
		src := sourcesQueue[0]

		sourcesQueue = sourcesQueue[1:]
		res = append(res, src)

		for _, dst := range g[src] {
			inDegree[dst]--
			if inDegree[dst] == 0 {
				sourcesQueue = append(sourcesQueue, dst)
			}
		}
	}

	if len(res) < vertices {
		return []int{}
	}

	return res
}

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
			return
		}
	}

	color[node] = black
	*topologicalOrder = append(*topologicalOrder, node)
}
