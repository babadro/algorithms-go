package _1857_largest_color_value_in_a_directed_graph

// doesn't work
func largestPathValue(colors string, edges [][]int) int {
	V := len(colors)
	g := New(V)
	for _, edge := range edges {
		g.AddEdge(edge[0], edge[1])
	}

	if g.IsCyclic() {
		return -1
	}

	max := 0
	f := func(path []int) {
		m := make(map[byte]int, len(colors))
		for _, v := range path {
			m[colors[v]]++
		}

		for _, count := range m {
			if count > max {
				max = count
			}
		}
	}

	g.AllPathsFromSourceToTarget(0, len(colors)-1, f)

	return max
}
