package _2092_find_all_people_with_secret

import (
	"github.com/babadro/algorithms-go/base/unionFind"
	"sort"
)

// tptl passed. best solution.
// https://leetcode.com/problems/find-all-people-with-secret/discuss/1601768/Union-Find-with-Disconnects
func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	union := make([]int, n)
	for i := range union {
		union[i] = -1
	}

	union[firstPerson] = 0

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})
	for i, j := 0, 0; i < len(meetings); i++ {
		if meetings[i][2] != meetings[j][2] {
			for k := j; k < i; k++ {
				p1, p2 := meetings[k][0], meetings[k][1]
				if find(union, p1) != 0 {
					union[p1] = -1
				}

				if find(union, p2) != 0 {
					union[p2] = -1
				}
			}

			j = i
		}

		root1, root2 := find(union, meetings[i][0]), find(union, meetings[i][1])
		if root1 != root2 {
			if root1 > root2 { // todo 1 need to understand
				root1, root2 = root2, root1
			}

			union[root2] = root1
		}
	}

	var res []int
	for i := range union {
		if find(union, i) == 0 {
			res = append(res, i)
		}
	}

	return res
}

func find(union []int, i int) int {
	val := union[i]
	for val != -1 {
		i = val
		val = union[i]
	}

	return i
}

// passed, but too long solution
func findAllPeople1(_ int, meetings [][]int, firstPerson int) []int {
	hasSecret := make(map[int]bool)
	hasSecret[0], hasSecret[firstPerson] = true, true

	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][2] < meetings[j][2]
	})

	var people, whoKnowsSecret []int
	for i := 0; i < len(meetings); {
		people, whoKnowsSecret = people[:0], whoKnowsSecret[:0]
		personToPeopleIDx := map[int]int{}
		j := i
		for ; j < len(meetings) && meetings[i][2] == meetings[j][2]; j++ {
			for k := 0; k < 2; k++ {
				person := meetings[j][k]
				if _, ok := personToPeopleIDx[person]; !ok {
					idx := len(personToPeopleIDx)
					personToPeopleIDx[person] = idx
					people = append(people, person)
				}
			}
		}

		u := unionFind.NewWQUPC(len(people))
		for k := i; k < j; k++ {
			person1, person2 := meetings[k][0], meetings[k][1]
			u.Union(personToPeopleIDx[person1], personToPeopleIDx[person2])
		}

		for _, p := range people {
			if hasSecret[p] {
				whoKnowsSecret = append(whoKnowsSecret, p)
			}
		}

		for _, p := range people {
			if !hasSecret[p] {
				for _, knowsSecret := range whoKnowsSecret {
					if u.Find(personToPeopleIDx[p], personToPeopleIDx[knowsSecret]) {
						hasSecret[p] = true

						break
					}
				}
			}
		}

		i = j
	}

	res := make([]int, 0, len(hasSecret))
	for person := range hasSecret {
		res = append(res, person)
	}

	return res
}
