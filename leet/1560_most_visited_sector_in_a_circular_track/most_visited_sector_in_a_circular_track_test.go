package _1560_most_visited_sector_in_a_circular_track

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_mostVisited(t *testing.T) {
	tests := []struct {
		n      int
		rounds []int
		want   []int
	}{
		{4, []int{1, 3, 1, 2}, []int{1, 2}},
		{2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}, []int{2}},
		{7, []int{1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
		{4, []int{4, 2}, []int{1, 2, 4}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("n=%d, rounds = %v", tt.n, tt.rounds), func(t *testing.T) {
			if got := mostVisited2(tt.n, tt.rounds); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostVisited() = %v, want %v", got, tt.want)
			}
		})
	}
}
