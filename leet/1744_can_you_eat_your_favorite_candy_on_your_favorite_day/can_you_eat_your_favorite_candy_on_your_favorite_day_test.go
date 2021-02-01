package _1744_can_you_eat_your_favorite_candy_on_your_favorite_day

import (
	"reflect"
	"testing"
)

func Test_canEat(t *testing.T) {
	tests := []struct {
		name         string
		candiesCount []int
		queries      [][]int
		want         []bool
	}{
		{
			"1",
			[]int{7, 4, 5, 3, 8}, [][]int{{0, 2, 2}, {4, 2, 4}, {2, 13, 100000000}},
			[]bool{true, false, true},
		},
		{
			"2",
			[]int{5, 2, 6, 4, 1}, [][]int{{3, 1, 2}, {4, 10, 3}, {3, 10, 100}, {4, 100, 30}, {1, 3, 1}},
			[]bool{false, true, true, false, false},
		},
		{
			"3", // todo 1 fails
			[]int{3, 4},
			[][]int{{1, 0, 7}},
			[]bool{true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canEat(tt.candiesCount, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("canEat() = %v, want %v", got, tt.want)
			}
		})
	}
}
