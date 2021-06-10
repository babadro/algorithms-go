package _1886_determine_whether_matrix_can_be_obtained_by_rotation

import "fmt"

// todo 1
func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	for y := 0; y < n; y++ {
		for x := 0; x < n; x++ {
			y1, y2 := y, n-1-y
			x1, x2 := x, n-1-x

			t := target[y][x]

			fmt.Println(y1, x1)
			fmt.Println(y2, x2)
			fmt.Println(x1, y1)
			fmt.Println(x2, y2)

			if mat[y1][x1] != t &&
				mat[y2][x2] != t &&
				mat[x1][y1] != t &&
				mat[x2][y2] != t {

			}
		}
	}

	return true
}
