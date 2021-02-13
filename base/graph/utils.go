package graph

func EdgesToAdjMatrix(vertexCount int, edges [][3]int) [][]int {
	matrix := make([][]int, vertexCount)
	for i := range matrix {
		matrix[i] = make([]int, vertexCount)
	}

	for _, edge := range edges {
		y, x, weight := edge[0], edge[1], edge[2]

		matrix[y][x] = weight
		matrix[x][y] = weight
	}

	return matrix
}
