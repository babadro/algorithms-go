package _2092_find_all_people_with_secret

import "sort"

// todo 1
func findAllPeople(_ int, meetings [][]int, firstPerson int) []int {
	schedule := make(map[int]map[int]bool)

	for _, meeting := range meetings {
		t := meeting[2]
		participants, ok := schedule[t]
		if !ok {
			participants = make(map[int]bool)
		}

		p1, p2 := meeting[0], meeting[1]
		participants[p1], participants[p2] = true, true

		schedule[t] = participants
	}

	hasSecret := make(map[int]bool)
	hasSecret[0], hasSecret[firstPerson] = true, true

	sortedMeetingTimes := make([]int, 0, len(schedule))
	for t := range schedule {
		sortedMeetingTimes = append(sortedMeetingTimes, t)
	}

	sort.Ints(sortedMeetingTimes)

	for _, t := range sortedMeetingTimes {
		participants := schedule[t]
		for participant := range participants {
			if hasSecret[participant] {
				for p := range participants {
					hasSecret[p] = true
				}

				break
			}
		}
	}

	res := make([]int, 0, len(hasSecret))
	for person := range hasSecret {
		res = append(res, person)
	}

	return res
}
