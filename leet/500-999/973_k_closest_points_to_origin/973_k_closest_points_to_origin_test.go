package _973_k_closest_points_to_origin

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_kClosest(t *testing.T) {
	tests := []struct {
		points [][]int
		k      int
		want   [][]int
	}{
		{[][]int{{1, 2}, {1, 3}}, 1, [][]int{{1, 2}}},
		{[][]int{{1, 3}, {3, 4}, {2, -1}}, 2, [][]int{{1, 3}, {2, -1}}},
		{[][]int{{4, 4}, {-3, -3}, {2, 2}, {1, 1}}, 2, [][]int{{1, 1}, {2, 2}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.points), func(t *testing.T) {
			got := kClosest(tt.points, tt.k)
			sort.Slice(got, func(i, j int) bool {
				return less(got, i, j)
			})
			sort.Slice(tt.want, func(i, j int) bool {
				return less(tt.want, i, j)
			})

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findClosestPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func less(points [][]int, i, j int) bool {
	xI, xJ := points[i][0], points[j][0]
	if xI != xJ {
		return xI < xJ
	}

	return points[i][1] < points[j][1]
}
