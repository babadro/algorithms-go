package _997_find_the_town_judge

// tptl. best solution
func findJudge(N int, trust [][]int) int {
	arr := make([]int, N+1)
	for _, pair := range trust {
		arr[pair[0]]--
		arr[pair[1]]++
	}

	for i := 1; i <= N; i++ {
		if arr[i] == N-1 {
			return i
		}
	}

	return -1
}

// passed. Not optimal
func findJudge2(N int, trust [][]int) int {
	graph := make([][2][]int, N)

	for _, pair := range trust {
		from := pair[0] - 1
		to := pair[1] - 1
		graph[from][0] = append(graph[from][0], to)
		graph[to][1] = append(graph[to][1], from)
	}

	for i := range graph {
		if len(graph[i][0]) == 0 && len(graph[i][1]) == N-1 {
			return i + 1
		}
	}

	return -1
}
