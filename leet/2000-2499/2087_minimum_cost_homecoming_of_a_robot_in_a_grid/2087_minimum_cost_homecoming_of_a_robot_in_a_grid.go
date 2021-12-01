package _2087_minimum_cost_homecoming_of_a_robot_in_a_grid

// passed. tptl. medium
func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	colChangeCost, rowChangeCost := 0, 0
	startCol, finishCol := startPos[1], homePos[1]
	startRow, finishRow := startPos[0], homePos[0]

	for col := startCol; col != finishCol; {
		if startCol < finishCol {
			col++
		} else {
			col--
		}

		colChangeCost += colCosts[col]
	}

	for row := startRow; row != finishRow; {
		if startRow < finishRow {
			row++
		} else {
			row--
		}

		rowChangeCost += rowCosts[row]
	}

	return colChangeCost + rowChangeCost
}
