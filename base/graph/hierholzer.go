package graph

// https://www.thecrazyprogrammer.com/2021/04/hierholzers-algorithm.html todo base
func EulerianCircuit(adj [][]int) []int {
	edges := make(map[int]int)

	for i := 0; i < len(adj); i++ {
		edges[i] = len(adj)
	}

	var currPath []int

	var circuit []int

	currPath = append(currPath, 0)

	currV := 0

	for len(currPath) != 0 {
		if edges[currV] > 0 {
			currPath
		}
	}
}
