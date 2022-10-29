package _2069_walking_robot_simulation_ii

const (
	east byte = iota
	north
	west
	south
)

type Robot struct {
	direction byte
	x, y      int
	w, h      int
}

func Constructor(width int, height int) Robot {
	return Robot{
		direction: east,
		x:         0,
		y:         height - 1,
		w:         width,
		h:         height,
	}
}

func (this *Robot) Step(num int) {
	for num > 0 {
		nextX, nextY := next(this.direction, this.x, this.y)
		if nextY == -1 || nextY == this.h || nextX == -1 || nextX == this.w {
			this.direction = (this.direction + 1) % 4
			continue
		}

		this.x, this.y = nextX, nextY

		num--
	}
}

func (this *Robot) GetPos() []int {
	return []int{this.x, this.h - 1 - this.y}
}

func (this *Robot) GetDir() string {
	switch this.direction {
	case north:
		return "North"
	case east:
		return "East"
	case south:
		return "South"
	case west:
		return "West"
	}

	return ""
}

func next(direction byte, x, y int) (int, int) {
	switch direction {
	case east:
		return x + 1, y
	case north:
		return x, y - 1
	case west:
		return x - 1, y
	case south:
		return x, y + 1
	}

	return -1, -1
}
