package _1351_count_negative_numbers_in_sorted_matrix

// tptl

// passed. simple
func countNegatives(grid [][]int) int {
	m, n := len(grid)-1, len(grid[0])-1

	res := 0
	for i := m; i >= 0; i-- {
		for j := n; j >= 0; j-- {
			if grid[i][j] < 0 {
				res++
			} else {
				break
			}
		}
	}

	return res
}

// passed. Binary search approach
func countNegatives2(grid [][]int) int {
	res := 0
	left := 0
	for i := len(grid) - 1; i >= 0; i-- {
		left = binarySearch(grid[i], left, len(grid[i])-1)
		if left == len(grid[i]) {
			break
		}
		res += len(grid[i]) - left
	}

	return res
}

func binarySearch(arr []int, left, right int) int {
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] >= 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}
