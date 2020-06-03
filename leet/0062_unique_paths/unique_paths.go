package _062_unique_paths

// Time limit exceeded
func uniquePathsRecursive(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	return uniquePaths(m-1, n) + uniquePaths(m, n-1)
}

func uniquePaths(m int, n int) int {
	matrix := make([][]int, m)
	for i := 0; i < m; i++ {
		matrix[i] = make([]int, n)
		matrix[i][0] = 1
	}
	for j := 0; j < n; j++ {
		matrix[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			matrix[i][j] = matrix[i-1][j] + matrix[i][j-1]
		}
	}

	return matrix[m-1][n-1]
}

// Time limit exceed
func uniquePathsByPerm(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	zeroCount, oneCount := m-1, n-1
	sequenceLen := zeroCount + oneCount
	arr := make([]byte, sequenceLen)
	for i := 0; i < oneCount; i++ {
		arr[i] = 1
	}
	counter := 0
	permCount(arr, &counter, 0, len(arr)-1)
	return counter
}

func permCount(nums []byte, res *int, l, r int) {
	if l == r {
		*res++
	} else {
		for i := l; i <= r; i++ {
			if shouldSwap(nums, l, i) {
				nums[i], nums[l] = nums[l], nums[i]
				permCount(nums, res, l+1, r)
				nums[i], nums[l] = nums[l], nums[i]
			}
		}
	}
}

// find ditinct permutations
func distinctPermutations(nums []byte, res *[][]byte, l, r int) {
	if l == r {
		permutation := make([]byte, len(nums))
		copy(permutation, nums)
		*res = append(*res, permutation)
	} else {
		for i := l; i <= r; i++ {
			if shouldSwap(nums, l, i) {
				nums[i], nums[l] = nums[l], nums[i]
				distinctPermutations(nums, res, l+1, r)
				nums[i], nums[l] = nums[l], nums[i]
			}
		}
	}
}

// TODO 2 need to understand
func shouldSwap(nums []byte, l, r int) bool {
	for i := l; i < r; i++ {
		if nums[i] == nums[r] {
			return false
		}
	}
	return true
}
