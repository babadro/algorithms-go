package _240_search_a_2d_matrix_ii

func searchMatrix(matrix [][]int, target int) bool {
	return false
}

func seach(xLeft, yLeft, xRight, yRight int, matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])

	if yLeft == yRight {
		return rowBinarySearch(matrix[yLeft][xLeft:xRight], target)
	}
	if xLeft == xRight {
		return columnBinarySearch(matrix, xLeft, yLeft, yRight, target)
	}

	for yLeft < m && xLeft < n && matrix[yLeft][xLeft] < target {
		yLeft++
		xLeft++
	}

	for yRight >= 0 && xRight >= 0 {
		yRight--
		xRight--
	}

	return false
}

func rowBinarySearch(nums []int, target int) bool {
	left, right := 0, len(nums)
	for left < right {
		middle := (right + left) / 2
		if nums[middle] > target {
			right = middle
			continue
		} else if nums[middle] < target {
			left = middle + 1
			continue
		}
		return true
	}
	return false
}

func columnBinarySearch(matrix [][]int, x, yFrom, yTo, target int) bool {
	for yFrom < yTo {
		middle := (yFrom + yTo) / 2
		if matrix[middle][x] > target {
			yTo = middle
			continue
		} else if matrix[middle][x] < target {
			yFrom = middle + 1
			continue
		}
		return true
	}
	return false
}

/* todo 2 this is not bad idea mabye...
func searchMatrix(matrix [][]int, target, lY, lX, rY, rX int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	leftY, leftX, rightY, rightX := 0, 0, len(matrix), len(matrix[0])

	for leftY < rightY && leftX < rightX {
		middleY, middleX := (leftY+rightY)/2, (leftX+rightX)/2
		if num := matrix[middleY][middleX]; num > target {
			rightX, rightY = middleX, middleY
			continue
		} else if num < target {
			leftX, leftY = middleX+1, middleY+1
			continue
		}
		return true
	}

}
*/

// brutforce. status: accepted
func searchMatrixBrutforce(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])

	for y := 0; y < m; y++ {
		for x := 0; x < n; x++ {
			if matrix[y][x] == target {
				return true
			}
		}
	}

	return false
}
