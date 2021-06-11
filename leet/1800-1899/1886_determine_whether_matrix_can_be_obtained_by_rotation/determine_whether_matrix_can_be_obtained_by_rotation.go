package _1886_determine_whether_matrix_can_be_obtained_by_rotation

// passed. easy, but tricky, so - tptl.
func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	a, b, c, d := true, true, true, true
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			m := mat[y][x]
			if a {
				a = m == target[y][x]
			}
			if b {
				b = m == target[n-1-y][n-1-x]
			}
			if c {
				c = m == target[x][n-1-y]
			}
			if d {
				d = m == target[n-1-x][y]
			}
		}
	}

	return a || b || c || d
}

// passed. mine.
func findRotation2(mat [][]int, target [][]int) bool {
	n := len(mat)
	f1 := func(y, x int) (int, int) {
		return y, x
	}
	f2 := func(y, x int) (int, int) {
		return n - 1 - y, n - 1 - x
	}
	f3 := func(y, x int) (int, int) {
		return x, n - 1 - y
	}
	f4 := func(y, x int) (int, int) {
		return n - 1 - x, y
	}

	return check(mat, target, f1) || check(mat, target, f2) || check(mat, target, f3) || check(mat, target, f4)
}

func check(mat, target [][]int, point func(int, int) (int, int)) bool {
	n := len(mat)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			y1, x1 := point(y, x)
			if mat[y][x] != target[y1][x1] {
				return false
			}
		}
	}

	return true
}
