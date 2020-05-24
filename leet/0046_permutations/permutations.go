package _046_permutations

// TODO 2 need to understand.
// try to find iterative solution
// complete next trask on "Sumbissions" page - Next Permutation, Permutations 2, Permutation Sequence, Combinations
func permute(nums []int) [][]int {
	var res [][]int
	perm(nums, &res, 0, len(nums)-1)
	return res
}

func perm(nums []int, res *[][]int, l, r int) {
	if l == r {
		permutation := make([]int, len(nums))
		copy(permutation, nums)
		*res = append(*res, permutation)
	} else {
		for i := l; i <= r; i++ {
			nums[l], nums[i] = nums[i], nums[l] // why should i do it when i == l (first iteration?) - скорее всего для краткости,
			// потому что иначе надо заключать в иф и эту строочку и обратный свап - громозко. Но работать будет.
			perm(nums, res, l+1, r)
			nums[l], nums[i] = nums[i], nums[l]
		}
	}
}
