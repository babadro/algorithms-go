package _2092_find_all_people_with_secret

import (
	"reflect"
	"sort"
	"testing"
)

func Test_findAllPeople(t *testing.T) {
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
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := findAllPeople(tt.n, tt.meetings, tt.firstPerson)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Log(tt.want)
				t.Log(got)
			}
		})
	}
}
