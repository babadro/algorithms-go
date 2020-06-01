package _062_unique_paths

// TODO 1
func uniquePaths(m int, n int) int {
	sequenceLen := m - 1 + n - 1
	arr := make([]byte, sequenceLen)

}

func perm(nums []byte, res *int, l, r int) {
	if l == r {
		*res++
	} else {
		for i := l; i <= r; i++ {
			nums[i], nums[l] = nums[l], nums[i]
			perm(nums, res, l+1, r)
			nums[i], nums[l] = nums[l], nums[i]
		}
	}
}
