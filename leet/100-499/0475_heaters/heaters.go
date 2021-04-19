package _475_heaters

import (
	"sort"
)

func abs(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// My solution ver 2.0
func findRadius(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	radius := 0
	for _, housePos := range houses {
		nearestHeaterPosIdx := sort.Search(len(heaters), func(i int) bool {
			return heaters[i] >= housePos
		})
		if nearestHeaterPosIdx == len(heaters) {
			nearestHeaterPosIdx--
		}
		minDistance := abs(housePos, heaters[nearestHeaterPosIdx])
		if nearestHeaterPosIdx > 0 {
			if distance := abs(heaters[nearestHeaterPosIdx-1], housePos); distance < minDistance {
				minDistance = distance
			}
		}
		if minDistance > radius {
			radius = minDistance
		}
	}
	return radius
}

// 40 ms solution
func findRadius40(houses []int, heaters []int) int {
	sort.Ints(houses)
	sort.Ints(heaters)
	w := 0
	m := 0
	for _, h := range houses {
		v := abs(h, heaters[w])
		for ; w < len(heaters)-1; w++ {
			vv := abs(h, heaters[w+1])
			if vv > v {
				break
			} else {
				v = vv
			}
		}
		if v > m {
			m = v
		}
	}
	return m
}

// My solution ver 1.0
func findRadius2(houses []int, heaters []int) int {
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
