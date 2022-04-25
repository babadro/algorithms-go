package _1_subsets

// tptl. leetcode 78
func findSubsetsIterative(nums []int) [][]int {
	res := [][]int{{}}
	for _, num := range nums {
		size := len(res)
		for i := 0; i < size; i++ {
			set := make([]int, len(res[i])+1)
			copy(set, res[i])
			set[len(set)-1] = num

			res = append(res, set)
		}
	}

	return res
}

func findSubsetsRecursive(nums []int) [][]int {
	res := [][]int{{}}
	rec(nums, 0, &res, nil)

	return res
}

func rec(nums []int, idx int, res *[][]int, cur []int) {
	if idx == len(nums) {
		return
	}

	rec(nums, idx+1, res, cur)

	cur = append(cur, nums[idx])
	*res = append(*res, cur)

	rec(nums, idx+1, res, cur)
}
