package _797_all_paths_from_source_to_target

// not effective
func allPathsSourceTarget2(graph [][]int) [][]int {
	return helper(graph, make([][][]int, len(graph)), 0)
}

func helper(graph [][]int, cache [][][]int, p int) [][]int {
	if p == len(graph)-1 {
		return [][]int{{p}}
	}

	if cache[p] != nil {
		return cache[p]
	}

	var res [][]int

	for _, v := range graph[p] {
		paths := helper(graph, cache, v)

		for _, childPath := range paths {
			path := make([]int, 0, len(childPath)+1)
			path = append(path, p)
			path = append(path, childPath...)

			res = append(res, path)
		}
	}

	cache[p] = res

	return res
}
