package _1964_find_the_longest_valid_obstacle_course_at_each_position

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_longestObstacleCourseAtEachPosition(t *testing.T) {
	tests := []struct {
		obstacles []int
		want      []int
	}{
		{[]int{1, 2, 3, 2}, []int{1, 2, 3, 3}},
		{[]int{2, 2, 1}, []int{1, 2, 1}},
		{[]int{3, 1, 5, 6, 4, 2}, []int{1, 1, 2, 3, 2, 2}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{tle1, tleOutput1},
	}
	for _, tt := range tests {
		name := fmt.Sprintf("%v", tt.obstacles)
		if len(tt.obstacles) > 10 {
			name = fmt.Sprintf("%v<...>", tt.obstacles[:10])
		}
		t.Run(name, func(t *testing.T) {
			if got := longestObstacleCourseAtEachPosition(tt.obstacles); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("longestObstacleCourseAtEachPosition() = %v, want %v", got, tt.want)
			}
		})
	}
}
