package _475_heaters

import "sort"

func findRadius(houses []int, heaters []int) int {
	n := 1_000_000_000
	return sort.Search(n, func(i int) bool {
		return checkRadius(houses, heaters, i)
	}) - 1
}

func checkRadius(houses []int, heaters []int, radius int) bool {
	heatersIdx := 0
	for _, housePos := range houses {
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
	}
	return true
}
