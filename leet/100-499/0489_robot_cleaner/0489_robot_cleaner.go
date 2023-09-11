package _489_robot_cleaner

// This is the robot's control interface.
// You should not implement it, or speculate about its implementation
type Robot struct {
}

// Returns true if the cell in front is open and robot moves into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool {
	return false
}

// Robot will stay in the same cell after calling TurnLeft/TurnRight.
// Each turn will be 90 degrees.
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}

// Clean the current cell.
func (robot *Robot) Clean() {}

var coords = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func cleanRoom(robot *Robot) {

}

var idx int

func walk(robot *Robot, coords [2]int, visited map[[2]int]bool) {
	if visited[coords] || !robot.Move() {
		rotate180()
		return
	}

	visited[coords] = true

	robot.Clean()

}

func rotate180() {
	idx %= idx + 2
}
