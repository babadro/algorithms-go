package _04_all_tasks_scheduling_orders

// todo 1
func allOrders(tasks int, prerequisites [][]int) [][]int {
	if tasks <= 0 {
		return nil
	}

	inDegree := make([]int, tasks)
	graph := make([][]int, tasks)

	for i := 0; i < len(prerequisites); i++ {
		parent := prerequisites[i][0]
		child := prerequisites[i][1]
		graph[parent] = append(graph[parent], child)
		inDegree[child]++
	}

	var sources []int
	for v, degree := range inDegree {
		if degree == 0 {
			sources = append(sources, v)
		}
	}

	var sortedOrder []int
	var res [][]int
	rec(graph, inDegree, sources, &sortedOrder, &res)

	return res
}

func rec(graph [][]int, inDegree []int, sources []int, sortedOrder *[]int, res *[][]int) {
	if len(sources) > 0 {
		for i, vertex := range sources {
			*sortedOrder = append(*sortedOrder, vertex)

			sourcesForNextCall := make([]int, len(sources)-1)
			copy(sourcesForNextCall, sources[:i])
			copy(sourcesForNextCall[i:], sources[i+1:])

			children := graph[vertex]
			for _, child := range children {
				inDegree[child]--
				if inDegree[child] == 0 {
					sourcesForNextCall = append(sourcesForNextCall, child)
				}
			}

			rec(graph, inDegree, sourcesForNextCall, sortedOrder, res)

			*sortedOrder = (*sortedOrder)[:len(*sortedOrder)-1]

			for _, child := range children {
				inDegree[child]++
			}
		}
	}

	// we don't know if the graph has a cycle (in this case topological sort doesn't exist)
	if len(*sortedOrder) == len(inDegree) {
		resArr := make([]int, len(*sortedOrder))
		copy(resArr, *sortedOrder)

		*res = append(*res, resArr)
	}
}
