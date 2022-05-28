package _046_permutations

// tptl. passed. iterative. easy to understand
func permute(nums []int) [][]int {
	var res [][]int
	q := [][]int{{}}
	for _, num := range nums {
		n := len(q)
		var oldPerm []int
		for i := 0; i < n; i++ {
			oldPerm, q = q[0], q[1:] // poll queue
			for j := 0; j <= len(oldPerm); j++ {
				newPerm := make([]int, len(oldPerm)+1)
				copy(newPerm[:j], oldPerm[:j])
				newPerm[j] = num
				copy(newPerm[j+1:], oldPerm[j:])

				if len(newPerm) == len(nums) {
					res = append(res, newPerm)
				} else {
					q = append(q, newPerm)
				}
			}
		}
	}

	return res
}

// TODO 2 need to understand.
// complete next task on "Sumbissions" page - Next Permutation, Permutations 2, Permutation Sequence, Combinations
func permute2(nums []int) [][]int {
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
