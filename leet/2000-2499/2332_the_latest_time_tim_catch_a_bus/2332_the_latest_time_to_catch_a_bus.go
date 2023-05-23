package _2332_the_latest_time_to_catch_a_bus

import "sort"

// #bnsrg tptl medium
func latestTimeCatchTheBus(buses []int, passengers []int, capacity int) int {
	sort.Ints(buses)
	sort.Ints(passengers)

	prev, res := -1, -1
	i := 0
	for _, busDepart := range buses {
		counter := 0
		for ; i < len(passengers) && passengers[i] <= busDepart && counter < capacity; i++ {
			p := passengers[i]

			if p-1 != prev && p-1 <= busDepart {
				res = p - 1
			}

			prev = p
			counter++
		}

		if counter == capacity {
			continue
		}

		if prev != busDepart {
			res = busDepart
		}
	}

	return res
}
