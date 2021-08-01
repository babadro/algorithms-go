package _1954_minimum_garden_perimeter_to_collect_enough_apples

// sum(155_362) = 10^15
const limit = int64(155_362)

func minimumPerimeter(neededApples int64) int64 {
	l, r := int64(1), limit
	for l < r {
		mid := (l + r) / 2
		if s := sum(mid); s < neededApples {
			l = mid + 1
		} else if s > neededApples {
			r = mid
		} else {
			return mid * 8
		}
	}

	return r * 8
}

func sum(radius int64) int64 {
	return 2 * radius * (radius + 1) * (2*radius + 1)
}
