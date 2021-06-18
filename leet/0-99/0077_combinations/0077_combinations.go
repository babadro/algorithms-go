package _0077_combinations

import "fmt"

// todo 1 need optimization
func combine(n int, k int) [][]int {
	curr := make([]int, k)
	res := make([][]int, 0)
	helper(n, 1, k, 0, curr, &res)

	return res
}

func helper(n, num, k, idx int, curr []int, res *[][]int) {
	if idx >= k {
		item := make([]int, k)
		copy(item, curr[:k])
		*res = append(*res, item)

		return
	}

	if num > n {
		return
	}

	fmt.Println(curr)

	helper(n, num+1, k, idx, curr, res)
	curr[idx] = num
	helper(n, num+1, k, idx+1, curr, res)
}
