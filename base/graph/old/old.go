package old

/*

// graph must be direct and acyclic (DAG)
// not effective solution.
func (g *Graph) AllPathsFromSourceToTarget(source, target int) [][]int {
	return g.allPathsHelper(source, target, make([][][]int, len(g.adj)))
}

func (g *Graph) allPathsHelper(source, target int, cache [][][]int) [][]int {
	if source == target {
		return [][]int{{source}}
	}

	if cache[source] != nil {
		return cache[source]
	}

	var res [][]int

	for _, v := range g.adj[source] {
		paths := g.allPathsHelper(v, target, cache)

		for _, childPath := range paths {
			path := make([]int, 0, len(childPath)+1)
			path = append(path, source)
			path = append(path, childPath...)

			res = append(res, path)
		}
	}

	cache[source] = res

	return res
}

// doesn't work
func (g *Graph) AllPathsFromSourceToTarget2(source, target int) [][]int {
	sortedVertexes := make([]int, 0, g.V())
	f := func(v int) {
		sortedVertexes = append(sortedVertexes, v)
	}

	g.TopologicalSort(f)

	d := make([][][]int, 0)

	n := g.V()
	for _, l := range sortedVertexes {
		if l == n {
			d[l] = [][]int{{n - 1}}
		} else {
			tmp := make([][]int, 0)
			for _, v := range g.adj[l] {
				for _, route := range d[v] {
					item := append([]int{l}, route...)
					tmp = append(tmp, item)
				}
				d[l] = tmp
			}
		}
	}

	return d[0]
}


// not efective solution
func (g *Graph) AllPathsFromSourceToTarget3(source, target int) [][]int {
	if g.V() == 0 || len(g.adj[source]) == 0 {
		return nil
	}

	memo := make([][][]int, g.V())

	return g.allPathsBFS(source, target, memo)
}

func (g *Graph) allPathsBFS(start, target int, memo [][][]int) [][]int {
	if memo[start] != nil {
		return memo[start]
	}

	var res [][]int
	if start == target {
		res = append(res, []int{start})
		memo[start] = res
		return res
	}

	for _, v := range g.adj[start] {
		if v == 0 {
			break
		}

		paths := g.allPathsBFS(v, target, memo)

		for _, path := range paths {
			arr := make([]int, 0, len(path)+1)
			arr = append(arr, start)
			arr = append(arr, path...)

			res = append(res, arr)
		}
	}

	memo[start] = res

	return res
}


// not effective solution
func (g *Graph) AllPathsFromSourceToTarget(source, target int) [][]int {
	sortedVertexes := make([]int, 0, g.V())
	f := func(v int) {
		sortedVertexes = append(sortedVertexes, v)
	}

	g.TopologicalSort(f)

	dp := make([]int, g.V())

	dp[target] = 1

	for i := len(sortedVertexes)-1; i >= 0; i-- {
		vertex := sortedVertexes[i]
		for j := 0; j < len(g.adj[vertex]); i-- {
			dp[vertex] += dp[g.adj[vertex][j]]
		}
	}

	return dp
}

// 8000 ns/op
func BenchmarkGraph_AllPathsFromSourceToTarget(b *testing.B) {
	var res [][]int

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
		g := Constructor(arr)
		b.StartTimer()

		res = g.AllPathsFromSourceToTarget(0, g.V()-1)
	}

	_ = res
}

//8000 ns/op
func BenchmarkGraph_AllPathsFromSourceToTarget3(b *testing.B) {
	var res [][]int

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		arr := [][]int{{3, 1}, {4, 6, 7, 2, 5}, {4, 6, 3}, {6, 4}, {7, 6, 5}, {6}, {7}, {}}
		g := Constructor(arr)
		b.StartTimer()

		res = g.AllPathsFromSourceToTarget3(0, g.V()-1)
	}

	_ = res
}

*/
