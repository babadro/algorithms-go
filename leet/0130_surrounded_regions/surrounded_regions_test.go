package _130_surrounded_regions

import "testing"

func TestSolve(t *testing.T) {
	matix1 := [][]byte{}
	solve(matix1)
	t.Log(matix1)

	matrix2 := [][]byte{{'X'}}
	solve(matrix2)
	t.Log(matrix2)

	matrix3 := [][]byte{{'X', 'O', 'O', 'X'}}
	solve(matrix3)
	t.Log(matrix3)

	matrix4 := [][]byte{
		{'X'},
		{'X'},
		{'O'},
	}
	solve(matrix4)
	t.Log(matrix4)

	matrix5 := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'X', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(matrix5)
	t.Log(matrix5)

	matrix6 := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(matrix6)
	t.Log(matrix6)

	matrix7 := [][]byte{
		{'X', 'X', 'X', 'X'},
		{'X', 'X', 'X', 'X'},
		{'X', 'O', 'O', 'X'},
		{'X', 'O', 'X', 'X'},
	}
	solve(matrix7)
	t.Log(matrix7)
}

func TestLeetFails(t *testing.T) {
	matrix := [][]byte{
		{'O', 'X', 'X', 'X', 'X', 'X', 'O', 'O'},
		{'O', 'O', 'O', 'X', 'X', 'X', 'X', 'O'},
		{'X', 'X', 'X', 'X', 'O', 'O', 'O', 'O'},
		{'X', 'O', 'X', 'O', 'O', 'X', 'X', 'X'},
		{'O', 'X', 'O', 'X', 'X', 'X', 'O', 'O'},
		{'O', 'X', 'X', 'O', 'O', 'X', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'X', 'X', 'O', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'O', 'X', 'X'},
	}

	expected := [][]byte{
		{'O', 'X', 'X', 'X', 'X', 'X', 'O', 'O'},
		{'O', 'O', 'O', 'X', 'X', 'X', 'X', 'O'},
		{'X', 'X', 'X', 'X', 'O', 'O', 'O', 'O'},
		{'X', 'O', 'X', 'X', 'O', 'X', 'X', 'X'},
		{'O', 'X', 'O', 'X', 'X', 'X', 'O', 'O'},
		{'O', 'X', 'X', 'O', 'O', 'X', 'X', 'O'},
		{'O', 'X', 'O', 'X', 'X', 'X', 'O', 'O'},
		{'O', 'X', 'X', 'X', 'X', 'O', 'X', 'X'},
	}

	solve(matrix)
	t.Log(matrix)
	t.Log(expected)
	w, h := len(matrix[0]), len(matrix)

	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			if matrix[y][x] != expected[y][x] {
				t.Error("mismatch y=", y, " x=", x)
			}
		}
	}

}
