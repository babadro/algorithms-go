package _40

// doesn't work todo 1
// https://leetcode.com/problems/largest-color-value-in-a-directed-graph/discuss/1198639/because-only-26-colors-C%2B%2B-O(n)
// todo 3 solution had made probably without 26 colors assumtion https://leetcode.com/problems/largest-color-value-in-a-directed-graph/discuss/1198765/Java-Simple-DFS-%2B-memo
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
