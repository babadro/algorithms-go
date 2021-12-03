package _2092_find_all_people_with_secret

import "sort"

// todo 1
func findAllPeople(_ int, meetings [][]int, firstPerson int) []int {
	hasSecret := make(map[int]bool)
	hasSecret[0], hasSecret[firstPerson] = true, true

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	for _, m := range meetings {
		p1, p2 := m[0], m[1]
		if hasSecret[p1] {
			hasSecret[p2] = true
		} else if hasSecret[p2] {
			hasSecret[p1] = true
		}
	}

	res := make([]int, 0, len(hasSecret))
	for person := range hasSecret {
		res = append(res, person)
	}

	return res
}
