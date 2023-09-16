package _489_robot_cleaner

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

var directions = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

var visited map[[2]int]bool

var matrix [][]int

var relativeX, relativeY, realX, realY, deltaY, deltaX, directionIDx int

type robotWrapper struct {
	r *Robot
}

func (r *robotWrapper) rotate180() {
	directionIDx = (directionIDx + 2) % 4
	r.r.TurnRight()
	r.r.TurnRight()
}

func (r *robotWrapper) turnLeft() {
	if directionIDx == 0 {
		directionIDx = 3
	} else {
		directionIDx--
	}

	r.r.TurnLeft()
}

func (r *robotWrapper) turnRight() {
	directionIDx = (directionIDx + 1) % 4

	r.r.TurnRight()
}

func (r *robotWrapper) move(checkVisited bool) bool {
	direction := directions[directionIDx]
	deltaY, deltaX = direction[0], direction[1]
	newY, newX := relativeY+deltaY, relativeX+deltaX

	key := [2]int{newY, newX}
	if (checkVisited && visited[key]) || !r.r.Move() {
		return false
	}

	if checkVisited {
		visited[key] = true
	}

	relativeY, relativeX = newY, newX

	return true
}

func (r *robotWrapper) clean() { r.r.Clean() }

func (r *robotWrapper) walk() {
	if r.move(true) {
		r.clean()
		r.walk()
	} else {
		r.rotate180()
	}

	// robot always returns from walk having opposite direction
	r.turnRight()
	// go right
	r.walk()

	// go left
	r.walk()

	// after returning take right direction to go back
	r.turnLeft()

	// go back
	r.move(false)
}

func cleanRoom(robot *Robot) {
	visited = make(map[[2]int]bool)
	relativeX, relativeY, directionIDx = 0, -1, 0

	r := robotWrapper{r: robot}

	r.walk()
	r.rotate180()
	r.move(false)
	r.rotate180()
	r.walk()
}
