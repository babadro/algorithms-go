package _475_heaters

import (
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	max, maxHeater := houses[len(houses)-1], heaters[len(heaters)-1]
	min, minHeater := houses[0], heaters[0]
	if maxHeater > max {
		max = maxHeater
	}
	if minHeater < min {
		min = minHeater
	}
	n := max - min + 1
	return sort.Search(n, func(i int) bool {
		return checkRadius(houses, heaters, i)
	})
}

func checkRadius(houses []int, heaters []int, radius int) bool {
	heatersIdx := 0
	for i := 0; i < len(houses); {
		housePos := houses[i]
		if heatersIdx == len(heaters) {
			return false
		}
		heatersPos := heaters[heatersIdx]
		if housePos > heatersPos+radius {
			heatersIdx++
			continue
		}
		if heatersPos-radius > housePos {
			return false
		}
		i++
	}
	return true
}
