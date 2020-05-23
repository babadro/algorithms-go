package _036_valid_sudoku

import "testing"

/*
func TestCoord(t *testing.T) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			t.Log((i*3)%9+j%3, ((i/3)*3)+j/3)
		}
	}
}
*/

func TestIsValidSudoku(t *testing.T) {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	t.Log(isValidSudoku(input))
}

func TestChangeArr(t *testing.T) {
	arr := make([]bool, 2)
	changeArr(arr, t)
	t.Log(arr[0], " ", arr[1])
}

func changeArr(arr []bool, t *testing.T) {
	arr[0], arr[1] = true, true
	t.Log(arr[0], arr[1])
}

var input = [][]byte{
	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
}

func BenchmarkMySol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidSudoku(input)
	}
}

func BenchmarkIsValidSudoku2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidSudoku2(input)
	}
}

func BenchmarkIsValidSudoku3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidSudoku3(input)
	}
}
func BenchmarkIsValidSudoku4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidSudoku4(input)
	}
}
