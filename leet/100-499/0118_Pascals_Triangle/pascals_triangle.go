package _118_Pascals_Triangle

func generate(numRows int) [][]int {
	res := make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		length := i + 1
		res[i] = make([]int, length)
		res[i][0], res[i][length-1] = 1, 1
		for j := 1; j < length-1; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}
