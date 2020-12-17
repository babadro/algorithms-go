package _874_walking_robot_simulation

const (
	north = 0
	east  = 1
	south = 2
	west  = 3

	left  = -2
	right = -1
)

// passed. todo 1: look for shorter solution
func robotSim(commands []int, obstacles [][]int) int {
	m := make(map[[2]int]bool)

	for _, point := range obstacles {
		m[[2]int{point[0], point[1]}] = true
	}

	curDirection := north
	x, y, max := 0, 0, 0
	for _, command := range commands {
		x, y, curDirection = nextState(command, curDirection, x, y, m)
		if dist := x*x + y*y; dist > max {
			max = dist
		}
	}

	return max
}

func turnLeft(currDirection int) int {
	if currDirection == north {
		return west
	}

	return currDirection - 1
}

func turnRight(currDirection int) int {
	if currDirection == west {
		return north
	}

	return currDirection + 1
}

func nextState(command, curDirect, curX, curY int, mObstacles map[[2]int]bool) (x, y, direct int) {
	if command == left {
		return curX, curY, turnLeft(curDirect)
	}

	if command == right {
		return curX, curY, turnRight(curDirect)
	}

	switch curDirect {
	case north:
		for command != 0 && !mObstacles[[2]int{curX, curY + 1}] {
			curY++
			command--
		}
	case east:
		for command != 0 && !mObstacles[[2]int{curX + 1, curY}] {
			curX++
			command--
		}
	case south:
		for command != 0 && !mObstacles[[2]int{curX, curY - 1}] {
			curY--
			command--
		}
	case west:
		for command != 0 && !mObstacles[[2]int{curX - 1, curY}] {
			curX--
			command--
		}
	}

	return curX, curY, curDirect
}
