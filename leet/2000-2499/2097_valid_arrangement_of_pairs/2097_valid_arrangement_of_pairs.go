package _2097_valid_arrangement_of_pairs

// todo 1
// https://www.geeksforgeeks.org/hierholzers-algorithm-directed-graph/
// https://www.thecrazyprogrammer.com/2021/04/hierholzers-algorithm.html
func validArrangement(pairs [][]int) [][]int {
	adj := make(map[int][]int)

	for _, pair := range pairs {
		adj[pair[0]] = append(adj[pair[0]], pair[1])
	}

	edges := make(map[int]int)
	for k, v := range adj {
		edges[k] = len(v)
	}

	currV := pairs[0][0]
	currPath := []int{currV}

	var circuit []int

	for len(currPath) > 0 {
		if edges[currV] != 0 {
			v := adj[currV][edges[currV]-1]
			currPath = append(currPath, v)

			edges[currV]--

			currV = currPath[len(currPath)-1]
		} else {
			last := len(currPath) - 1
			circuit = append(circuit, currPath[last])

			currPath = currPath[:last]
		}
	}

	res := make([][]int, 0, len(pairs))
	j := 0
	var pair []int
	for i := len(circuit) - 1; i >= 0; {
		if j%2 == 0 {
			pair = make([]int, 2)
			pair[0] = circuit[i]

			i--
		} else {
			pair[1] = circuit[i]

			res = append(res, pair)
		}

		j++
	}

	return res
}
