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

var directions = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

var visited map[[2]int]bool

var matrix [][]int

type robotWrapper struct {
	x, y         int
	directionIDx int
	r            *Robot
}

func (r *robotWrapper) rotate180() {
	r.directionIDx = (r.directionIDx + 2) % 4
	r.r.TurnRight()
	r.r.TurnRight()
}

func (r *robotWrapper) turnLeft() {
	if r.directionIDx == 0 {
		r.directionIDx = 3
	} else {
		r.directionIDx--
	}

	r.r.TurnLeft()
}

func (r *robotWrapper) turnRight() {
	r.directionIDx = (r.directionIDx + 1) % 4

	r.r.TurnRight()
}

func (r *robotWrapper) move(checkVisited bool) bool {
	direction := directions[r.directionIDx]
	deltaY, deltaX := direction[0], direction[1]
	newY, newX := r.y+deltaY, r.x+deltaX

	key := [2]int{newY, newX}
	if (checkVisited && visited[key]) || !r.r.Move() {
		return false
	}

	if checkVisited {
		visited[key] = true
	}

	r.y, r.x = newY, newX

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

	r := robotWrapper{r: robot}

	r.walk()
	r.rotate180()
	r.move(false)
	r.rotate180()
	r.walk()
}
