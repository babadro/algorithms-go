package _036_valid_sudoku

import "testing"

func TestCoord(t *testing.T) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			t.Log((i*3)%9+j%3, (i*3)%9+j/3)
		}
	}
}
