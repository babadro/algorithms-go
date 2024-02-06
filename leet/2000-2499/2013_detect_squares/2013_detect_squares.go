package _2013_detect_squares

type coords struct {
	y, x int
}

// #bnsrg medium passed
type DetectSquares struct {
	points map[coords]int
	yToXes map[int][]int
}

func Constructor() DetectSquares {
	return DetectSquares{
		points: make(map[coords]int),
		yToXes: make(map[int][]int),
	}
}

func (this *DetectSquares) Add(point []int) {
	p := coords{x: point[0], y: point[1]}

	if this.points[p] == 0 {
		this.yToXes[p.y] = append(this.yToXes[p.y], p.x)
	}

	this.points[p]++
}

func (this *DetectSquares) Count(point []int) int {
	p := coords{x: point[0], y: point[1]}

	res := 0

	for _, x := range this.yToXes[p.y] {
		if x == p.x {
			continue
		}

		edgeLen := p.x - x
		if edgeLen < 0 {
			edgeLen *= -1
		}

		p2, p3, p4 := coords{y: p.y, x: x}, coords{y: p.y + edgeLen, x: x}, coords{y: p.y + edgeLen, x: p.x}

		res += this.points[p2] * this.points[p3] * this.points[p4]

		p2, p3, p4 = coords{y: p.y, x: x}, coords{y: p.y - edgeLen, x: x}, coords{y: p.y - edgeLen, x: p.x}

		res += this.points[p2] * this.points[p3] * this.points[p4]
	}

	return res
}
