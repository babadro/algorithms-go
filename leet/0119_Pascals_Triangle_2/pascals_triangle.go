package _119_Pascals_Triangle_2

func getRow(rowIndex int) []int {
	length := rowIndex + 1
	triangle := make([][]int, length)
	for i := 0; i < length; i++ {
		length := i + 1
		triangle[i] = make([]int, length)
		triangle[i][0], triangle[i][length-1] = 1, 1
		for j := 1; j < length-1; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}
	return triangle[rowIndex]
}
