package _874_walking_robot_simulation

const (
	up    = 0
	left  = 1
	right = 2
	down  = 3
)

// todo 1
func robotSim(commands []int, obstacles [][]int) int {
	m := make(map[[2]int]bool)

	for _, point := range obstacles {
		m[[2]int{point[0], point[1]}] = true
	}

}

func changeDirection(command, currDirection int) int {

}
