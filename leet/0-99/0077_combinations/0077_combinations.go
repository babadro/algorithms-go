package _0077_combinations

// passed. best solution. tptl. medium.
func combine(n int, k int) [][]int {
	curr := make([]int, 0, k)
	res := make([][]int, 0)
	helper(n, 1, k, k, curr, &res)

	return res
}

func helper(n, num, k, size int, curr []int, res *[][]int) {
	if size == 0 {
		item := make([]int, len(curr))
		copy(item, curr)
		*res = append(*res, item)

		return
	}

	if n+1-num < size {
		return
	}

	helper(n, num+1, k, size, curr, res)
	curr = append(curr, num)
	helper(n, num+1, k, size-1, curr, res)
}
