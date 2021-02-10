package unionFind

func find(id []int, i int) int {
	if id[i] == -1 {
		return i
	}

	return find(id, id[i])
}

func Union(id []int, x, y int) {
	xSet := find(id, x)
	ySet := find(id, y)
	id[xSet] = ySet
}

func IsCycle(itemsCount int, pairs [][2]int) bool {
	id := make([]int, itemsCount)

	for i := 0; i < itemsCount; i++ {
		id[i] = -1
	}

	for i := 0; i < len(pairs); i++ {
		x := find(id, pairs[i][0])
		y := find(id, pairs[i][1])

		if x == y {
			return true
		}

		Union(id, x, y)
	}

	return false
}
