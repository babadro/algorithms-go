package _901_find_a_peak_element_ii

// passed. todo2 binary search solution is faster
func findPeakGrid(mat [][]int) []int {
	for y := range mat {
		for x := range mat[y] {
			el := mat[y][x]
			if x > 0 && el < mat[y][x-1] ||
				x < len(mat[y])-1 && el < mat[y][x+1] ||
				y > 0 && el < mat[y-1][x] ||
				y < len(mat)-1 && el < mat[y+1][x] {
				continue
			}

			return []int{y, x}
		}
	}

	return nil
}
