package _1833_maximum_ice_cream_bars

import "sort"

// passed. easy todo 3 make it with minHeap - it's more effective
func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)

	count := 0
	for _, cost := range costs {
		if cost <= coins {
			count++
			coins -= cost
		} else {
			break
		}
	}

	return count
}
