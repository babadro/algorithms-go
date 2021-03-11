package graph

func EdgesToAdjMatrix(vertexCount int, edges [][3]int, _1_indexed bool) [][]int {
	matrix := make([][]int, vertexCount)
	for i := range matrix {
		matrix[i] = make([]int, vertexCount)
	}

	for _, edge := range edges {
		var y, x int
		if _1_indexed {
			y, x = edge[0]-1, edge[1]-1
		} else {
			y, x = edge[0], edge[1]
		}

		weight := edge[2]

		matrix[y][x] = weight
		matrix[x][y] = weight
	}

	return matrix
}

func EdgesToAdjList(vertexCount int, edges [][3]int, _1_indexed bool) [][][2]int {
	adjList := make([][][2]int, vertexCount)

	for _, edge := range edges {
		var y, x int
		if _1_indexed {
			y, x = edge[0]-1, edge[1]-1
		} else {
			y, x = edge[0], edge[1]
		}

		weight := edge[2]

		adjList[y] = append(adjList[y], [2]int{weight, x})
		adjList[x] = append(adjList[x], [2]int{weight, y})
	}

	return adjList
}

func EdgesSliceToAdjList(vertexCount int, edges [][]int, _1_indexed bool) [][][2]int {
	adjList := make([][][2]int, vertexCount)

	for _, edge := range edges {
		var y, x int
		if _1_indexed {
			y, x = edge[0]-1, edge[1]-1
		} else {
			y, x = edge[0], edge[1]
		}

		weight := edge[2]

		adjList[y] = append(adjList[y], [2]int{weight, x})
		adjList[x] = append(adjList[x], [2]int{weight, y})
	}

	return adjList
}
