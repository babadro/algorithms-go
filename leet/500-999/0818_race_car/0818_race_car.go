package _0818_race_car

// #bnsrg
// todo 2 find shorter and faster (probably dp) solution
func racecar(target int) int {
	visited := make(map[[2]int]bool)
	var queue [][3]int
	item := [3]int{0, 1, 0} // position, speed, movies

	queue = append(queue, item)
	for len(queue) > 0 {
		item, queue = queue[0], queue[1:]

		position, speed, movies := item[0], item[1], item[2]

		if visited[[2]int{position, speed}] {
			continue
		}

		if position == target {
			return movies
		}

		visited[[2]int{position, speed}] = true

		queue = append(queue, [3]int{position + speed, speed * 2, movies + 1})

		if position+speed > target && speed > 0 {
			speed = -1

			queue = append(queue, [3]int{position, speed, movies + 1})
		} else if position+speed < target && speed < 0 {
			speed = 1

			queue = append(queue, [3]int{position, speed, movies + 1})
		}
	}

	return -1
}
