package _200_number_of_islands

import "testing"

func TestNumOfIslands(t *testing.T) {
	input1 := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	input2 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	input3 := [][]byte{
		{'1', '1', '0', '0', '1'},
		{'1', '1', '1', '1', '1'},
		{'0', '0', '0', '0', '0'},
		{'1', '0', '0', '0', '0'},
	}

	input4 := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}

	// TODO failed test case
	input5 := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}

	t.Log(numIslands(input1))
	t.Log(numIslands(input2))
	t.Log(numIslands(input3))
	t.Log(numIslands(input4))
	t.Log(numIslands(input5))
}

/*

 */
