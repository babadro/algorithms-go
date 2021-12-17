package graph

// undirected graph
// https://www.techiedelight.com/eulerian-path-undirected-graph/ has eulerian path or cycle todo base
func UndirectedGraphHasEulerianPathOrCycle(g Graph) (hasPath, hasCycle bool) {
	if !IsConnected(g) {
		return false, false
	}

	var oddCount int
	for i := range g.adj {
		if len(g.adj[i])%2 == 1 {
			oddCount++
		}
	}

	if oddCount == 0 {
		return true, true
	}

	if oddCount == 2 {
		return true, false
	}

	return false, false
}
