package _489_robot_cleaner

var realX, realY int

var matrix [][]int

// This is the robot's control interface.
// You should not implement it, or speculate about its implementation
type Robot struct {
}

// Returns true if the cell in front is open and robot moves into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool {
	newX := realX + deltaX
	if newX < 0 || newX == len(matrix[0]) {
		return false
	}

	newY := realY + deltaY
	if newY < 0 || newY == len(matrix) {
		return false
	}

	if matrix[newY][newX] == 0 {
		return false
	}

	realX, realY = newX, newY

	return true
}

// Robot will stay in the same cell after calling TurnLeft/TurnRight.
// Each turn will be 90 degrees.
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}

// Clean the current cell.
func (robot *Robot) Clean() {
	matrix[realY][realX] = 2
}
