package _406_queue_reconstruction_by_height

import "sort"

// best solution. tptl. passed. not mine
func reconstructQueue2(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		return people[i][0] > people[j][0]
	})

	ret := make([][]int, 0)
	for _, p := range people {
		ret = append(ret, []int{})
		copy(ret[p[1]+1:], ret[p[1]:])
		ret[p[1]] = p
	}
	return ret
}

// passed. bruteforce
func reconstructQueue(people [][]int) [][]int {
	n := len(people)
	res := make([][]int, n)

	sort.Slice(people, func(i, j int) bool {
		l, r := people[i], people[j]
		if l[0] == r[0] {
			return l[1] < r[1]
		}

		return l[0] < r[0]
	})

	for _, p := range people {
		i, counter, height, count := 0, 0, p[0], p[1]
		for ; counter < count || (i < n && res[i] != nil); i++ {
			if res[i] == nil || res[i][0] >= height {
				counter++
			}
		}

		res[i] = p
	}

	return res
}

// doesn't work
func reconstructQueueOld(people [][]int) [][]int {
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
