package _797_all_paths_from_source_to_target

// tptl. best solution. passed.
func allPathsSourceTarget(graph [][]int) [][]int {
	tmp, res := make([]int, 0), make([][]int, 0)
	traverse(graph, 0, tmp, &res)

	return res
}

func traverse(graph [][]int, V int, tmp []int, res *[][]int) {
	tmp = append(tmp, V)

	if V == len(graph)-1 {
		arr := make([]int, len(tmp))
		copy(arr, tmp)
		*res = append(*res, arr)
		return
	}

	for _, v := range graph[V] {
		traverse(graph, v, tmp, res)
	}
}
