package _939_miminum_area_rectangle

import (
	"math"
	"slices"
	"sort"
)

// #bnsrg passed medium
// todo 2 shorter and faster solution exists
func minAreaRect(points [][]int) int {
	xPoints, yPoints, allPoints := make(map[int][]int), make(map[int][]int), make(map[[2]int]bool)

	for _, p := range points {
		x, y := p[0], p[1]

		yPoints[x] = append(yPoints[x], y)
		xPoints[y] = append(xPoints[y], x)

		allPoints[[2]int{x, y}] = true
	}

	for _, p := range xPoints {
		sort.Ints(p)
	}

	for _, p := range yPoints {
		sort.Ints(p)
	}

	minArea := math.MaxInt64

	for x, yArr := range yPoints {
		for i := 0; i < len(yArr)-1; i++ {
			y := yArr[i]

			for j := i + 1; j < len(yArr); j++ {
				y2 := yArr[j]

				xArr := xPoints[y2]

				idx, found := slices.BinarySearch(xArr, x)
				if !found {
					panic("impossible")
				}

				for k := idx + 1; k < len(xArr); k++ {
					x2 := xArr[k]

					area := (y2 - y) * (x2 - x)

					if area >= minArea {
						break
					}

					if allPoints[[2]int{x2, y}] {
						minArea = area

						break
					}
				}
			}
		}
	}

	if minArea == math.MaxInt64 {
		return 0
	}

	return minArea
}
