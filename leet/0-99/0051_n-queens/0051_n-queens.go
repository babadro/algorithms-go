package _051_n_queens

// tptl. passed. hard. fast.
func solveNQueens(n int) [][]string {
	// initial state is [0, 1, 2, 3, ..., n-1]
	// each element in the array is the position Queen in the row.
	state := make([]int, n)
	for i := range state {
		state[i] = i
	}

	var res [][]string
	// generate permutations
	perm(state, &res, 0, len(state)-1)

	return res
}

func perm(state []int, res *[][]string, l, r int) {
	if l == r {
		if checkDiagonal(state, l) {
			*res = append(*res, convertPermToSolution(state))
		}
	} else {
		for i := l; i <= r; i++ {
			state[l], state[i] = state[i], state[l]

			if checkDiagonal(state, l) {
				perm(state, res, l+1, r)
			}

			state[l], state[i] = state[i], state[l]
		}
	}
}

func convertPermToSolution(perm []int) []string {
	res := make([]string, len(perm))
	for i := range res {
		b := make([]byte, len(perm))
		for j := range b {
			b[j] = '.'
		}

		b[perm[i]] = 'Q'

		res[i] = string(b)
	}

	return res
}

func checkDiagonal(perm []int, currIDx int) bool {
	cur := perm[currIDx]
	for i := currIDx - 1; i >= 0; i-- {
		prev, diff := perm[i], currIDx-i

		// the queen on one of the previous rows occupies the same diagonal
		if prev == cur-diff || prev == cur+diff {
			return false
		}
	}

	return true
}
