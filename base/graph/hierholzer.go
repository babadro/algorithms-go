package graph

// https://www.thecrazyprogrammer.com/2021/04/hierholzers-algorithm.html
// https://www.geeksforgeeks.org/hierholzers-algorithm-directed-graph/
// dyx
func EulerianCircuit(adj [][]int) []int {
	curPath, circuit := []int{0}, make([]int, 0, len(adj))

	for len(curPath) > 0 {
		cur := curPath[len(curPath)-1]

		if len(adj[cur]) > 0 {
			last := len(adj[cur]) - 1

			curPath = append(curPath, adj[cur][last])
			adj[cur] = adj[cur][:last]
		} else {
			last := len(curPath) - 1

			circuit = append(circuit, curPath[last])
			curPath = curPath[:last]
		}
	}

	last := len(circuit) - 1
	for i := 0; i < len(circuit)/2; i++ {
		circuit[i], circuit[last-i] = circuit[last-i], circuit[i]
	}

	return circuit
}

// https://leetcode.com/problems/valid-arrangement-of-pairs/
func EulerianPathOrCircuit(adj [][]int) []int {
	cnt := make(map[int]int)
	for from := range adj {
		for _, to := range adj[from] {
			cnt[from]++
			cnt[to]--
		}
	}

	start := adj[0][0]
	for vertex, inOutDegree := range cnt {
		if inOutDegree > 0 { // eulerian path
			start = vertex
			break
		}
	}

	stack, path := []int{start}, make([]int, 0, len(adj))

	for len(stack) > 0 {
		cur := stack[len(stack)-1]

		if len(adj[cur]) > 0 {
			last := len(adj[cur]) - 1

			stack = append(stack, adj[cur][last])
			adj[cur] = adj[cur][:last]
		} else {
			last := len(stack) - 1

			path = append(path, stack[last])
			stack = stack[:last]
		}
	}

	last := len(path) - 1
	for i := 0; i < len(path)/2; i++ {
		path[i], path[last-i] = path[last-i], path[i]
	}

	return path
}

// another yet variant.
func eulerianCircuit2(adj [][]int) []int {
	stack, route := []int{0}, make([]int, 0, len(adj))

	for len(stack) > 0 {
		for len(adj[stack[len(stack)-1]]) > 0 {
			i := stack[len(stack)-1]

			last := len(adj[i]) - 1
			stack = append(stack, adj[i][last])
			adj[i] = adj[i][:last]
		}

		last := len(stack) - 1
		route = append(route, stack[last])
		stack = stack[:last]
	}

	last := len(route) - 1
	for i := 0; i < len(route)/2; i++ {
		route[i], route[last-i] = route[last-i], route[i]
	}

	return route
}
