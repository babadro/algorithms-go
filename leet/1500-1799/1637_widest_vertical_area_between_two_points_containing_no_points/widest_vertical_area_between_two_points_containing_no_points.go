package _637_widest_vertical_area_between_two_points_containing_no_points

import "sort"

// passed. tptl
func maxWidthOfVerticalArea(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	max := 0
	for i := 0; i < len(points)-1; i++ {
		if diff := points[i+1][0] - points[i][0]; diff > max {
			max = diff
		}
	}

	return max
}
