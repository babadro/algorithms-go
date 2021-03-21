package _791_find_center_of_star_graph

// passed. easy
func findCenter(edges [][]int) int {
	m := make(map[int]bool)

	for _, edge := range edges {
		v1, v2 := edge[0], edge[1]

		if m[v1] {
			return v1
		}
		m[v1] = true

		if m[v2] {
			return v2
		}

		m[v2] = true
	}

	return 0
}
