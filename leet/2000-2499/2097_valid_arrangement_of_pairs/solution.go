package _2097_valid_arrangement_of_pairs

// passed dyx hard #graph
func validArrangement(pairs [][]int) [][]int {
	cnt := make(map[int]int)
	graph := make(map[int][]int)
	for _, p := range pairs {
		graph[p[0]] = append(graph[p[0]], p[1])
		cnt[p[0]]++
		cnt[p[1]]--
	}

	start := pairs[0][0]
	for vertex, inOutDegree := range cnt {
		if inOutDegree > 0 {
			start = vertex
			break
		}
	}

	stack, path := []int{start}, make([]int, 0, len(graph))

	for len(stack) > 0 {
		cur := stack[len(stack)-1]

		if len(graph[cur]) > 0 {
			last := len(graph[cur]) - 1

			stack = append(stack, graph[cur][last])
			graph[cur] = graph[cur][:last]
		} else {
			last := len(stack) - 1

			path = append(path, stack[last])
			stack = stack[:last]
		}
	}

	res := make([][]int, 0, len(graph))
	for i := len(path) - 1; i > 0; i-- {
		res = append(res, []int{path[i], path[i-1]})
	}

	return res
}
