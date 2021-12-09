package _2097_valid_arrangement_of_pairs

// todo 1
// https://www.geeksforgeeks.org/hierholzers-algorithm-directed-graph/
// https://www.thecrazyprogrammer.com/2021/04/hierholzers-algorithm.html
func validArrangement(pairs [][]int) [][]int {
	adj := make(map[int][]int)

	for _, pair := range pairs {
		adj[pair[0]] = append(adj[pair[0]], pair[1])
	}

	currPath := []int{pairs[0][0]}

	var circuit []int

	for len(currPath) > 0 {

		lastCurrPath := len(currPath) - 1
		currV := currPath[lastCurrPath]

		if nodes := adj[currV]; len(nodes) != 0 {
			last := len(nodes) - 1
			nextV := adj[currV][last]
			currPath = append(currPath, nextV)

			adj[currV] = adj[currV][:last]
		} else {
			circuit = append(circuit, currPath[lastCurrPath])

			currPath = currPath[:lastCurrPath]
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

func eulerianPath(adj map[int][]int, start int) [][]int {
	currPath := []int{start}

	var path []int

	for len(currPath) > 0 {

		lastCurrPath := len(currPath) - 1
		currV := currPath[lastCurrPath]

		if nodes := adj[currV]; len(nodes) != 0 {
			last := len(nodes) - 1
			nextV := adj[currV][last]
			currPath = append(currPath, nextV)

			adj[currV] = adj[currV][:last]
		} else {
			path = append(path, currPath[lastCurrPath])

			currPath = currPath[:lastCurrPath]
		}
	}

	res := make([][]int, 0, len(path)-1)
	j := 0
	var pair []int
	for i := len(path) - 1; i >= 0; {
		if j%2 == 0 {
			pair = make([]int, 2)
			pair[0] = path[i]

			i--
		} else {
			pair[1] = path[i]

			res = append(res, pair)
		}

		j++
	}

	return res
}
