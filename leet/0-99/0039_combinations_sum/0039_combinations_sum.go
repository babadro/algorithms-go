package _039_combinations_sum

import "sort"

// passed. medium
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	req(candidates, target, &[]int{}, 0, &res)

	return res
}

func req(arr []int, target int, cur *[]int, curLen int, res *[][]int) {
	if target == 0 {
		newArr := make([]int, len(*cur))
		copy(newArr, *cur)
		*res = append(*res, newArr)

		return
	}

	if target < 0 || len(arr) == 0 {
		return
	}

	req(arr[1:], target, cur, curLen, res)
	*cur = append((*cur)[:curLen], arr[0])
	req(arr, target-arr[0], cur, curLen+1, res)
}
