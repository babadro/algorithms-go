package _2092_find_all_people_with_secret

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findAllPeople(t *testing.T) {
	// [][]int{{0, 1, 1}, {0, 2, 1}, {0, 3, 1}, {0, 4, 1}...., {0, 99_999, 1}}
	tl1Input := make([][]int, 99_999)
	for i := range tl1Input {
		tl1Input[i] = []int{0: 0, 1: i + 1, 2: 1}
	}

	tle1Want := make([]int, 100_000)
	for i := range tle1Want {
		tle1Want[i] = i
	}

	tests := []struct {
		n           int
		meetings    [][]int
		firstPerson int
		want        []int
	}{
		{6, [][]int{{1, 2, 5}, {2, 3, 8}, {1, 5, 10}}, 1, []int{0, 1, 2, 3, 5}},
		{4, [][]int{{3, 1, 3}, {1, 2, 2}, {0, 3, 3}}, 3, []int{0, 1, 3}},
		{5, [][]int{{3, 4, 2}, {1, 2, 1}, {2, 3, 1}}, 1, []int{0, 1, 2, 3, 4}},
		{6, [][]int{{0, 2, 1}, {1, 3, 1}, {4, 5, 1}}, 1, []int{0, 1, 2, 3}},
		{7336, bigInput1, 6384, expected1},
		{100_000, tl1Input, 1, tle1Want},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := findAllPeople(tt.n, tt.meetings, tt.firstPerson)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Error(got)
				t.Log(tt.want)
			}
		})
	}
}
