package _489_robot_cleaner

// bnsrg #hard passed
// todo 2 find shorter solution - there are a lot of them
func cleanRoom2(robot *Robot) {
	visited = make(map[[2]int]bool)
	relativeX, relativeY, directionIDx = 0, 0, 0

	r := robotWrapper{r: robot}
	r.clean()

	r.walk()
	r.turnRight()

	r.walk()
	r.turnRight()

	r.walk()
	r.turnRight()

	r.walk()
}

var directions = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

var visited map[[2]int]bool

var relativeX, relativeY, deltaY, deltaX, directionIDx int

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

// robot always returns to the same cell after walk having opposite direction
func (r *robotWrapper) walk() {
	if r.move(true) {
		r.clean()
		r.walk()
	} else {
		r.rotate180()
		return
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
