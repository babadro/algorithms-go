package _43

import (
	"fmt"
	"testing"
)

func Test_minSkips(t *testing.T) {
	tests := []struct {
		dist        []int
		speed       int
		hoursBefore int
		want        int
	}{
		{[]int{1, 3, 2}, 4, 2, 1},
		{[]int{7, 3, 5, 5}, 2, 10, 2},
		{[]int{7, 3, 5, 5}, 1, 10, -1},
		{[]int{2, 1, 5, 4, 4, 3, 2, 9, 2, 10}, 6, 7, 7},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.dist), func(t *testing.T) {
			if got := minSkips(tt.dist, tt.speed, tt.hoursBefore); got != tt.want {
				t.Errorf("minSkips() = %v, want %v", got, tt.want)
			}
		})
	}
}
