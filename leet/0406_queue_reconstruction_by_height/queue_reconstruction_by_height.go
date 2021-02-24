package _406_queue_reconstruction_by_height

import "sort"

// todo 1
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		l, r := people[i], people[j]
		sumL, sumR := l[0]+l[1], r[0]+r[1]
		if sumL == sumR {
			return l[0] > r[0]
		}

		return sumL < sumR
	})

	res := make([][]int, 0, len(people))

Loop:
	for len(people) != 0 {
		p := people[0]
		height, count := p[0], p[1]
		counter := 0
		for i := len(res) - 1; i >= 0; i-- {
			if res[i][0] >= height {
				counter++
			}
		}

		if counter < count {
			people[0], people[1] = people[1], people[0]
			continue Loop
		}

		res = append(res, p)
		people = people[1:]

		if counter > count {
			last := len(res) - 1
			res[last], res[last-1] = res[last-1], res[last]
		}
	}

	return res
}
