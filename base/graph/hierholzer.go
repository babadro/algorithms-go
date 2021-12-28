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
