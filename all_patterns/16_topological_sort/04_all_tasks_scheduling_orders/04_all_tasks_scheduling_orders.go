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
	return getAllTopologicalSorts(graph, inDegree, sources, &sortedOrder)
}

func getAllTopologicalSorts(graph [][]int, inDegree []int, sources []int, sortedOrder *[]int) [][]int {
	if len(sources) > 0 {
		for _, vertex := range sources {
			*sortedOrder = append(*sortedOrder, vertex)
			sourcesForNextCall := make([]int, len(sources))
			copy(sourcesForNextCall, sources)
		}
	}
}
