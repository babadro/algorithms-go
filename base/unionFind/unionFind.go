package unionFind

func find(parent []int, i int) int {
	if parent[i] == -1 {
		return i
	}

	return find(parent, parent[i])
}

func Union(parent []int, x, y int) {
	xSet := find(parent, x)
	ySet := find(parent, y)
	parent[xSet] = ySet
}

// todo 1 understand
func IsCycle(itemsCount int, pairs [][2]int) bool {
	parent := make([]int, itemsCount)

	for i := 0; i < itemsCount; i++ {
		parent[i] = -1
	}

	for i := 0; i < len(pairs); i++ {
		x := find(parent, pairs[i][0])
		y := find(parent, pairs[i][1])

		if x == y {
			return true
		}

		Union(parent, x, y)
	}

	return false
}
