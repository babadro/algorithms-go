package _37

import (
	"reflect"
	"testing"
)

func Test_getOrder(t *testing.T) {
	tests := []struct {
		name  string
		tasks [][]int
		want  []int
	}{
		{tasks: [][]int{{1, 2}, {2, 4}, {3, 2}, {4, 1}}, want: []int{0, 2, 3, 1}},
		{tasks: [][]int{{7, 10}, {7, 12}, {7, 5}, {7, 4}, {7, 2}}, want: []int{4, 3, 2, 0, 1}},
		{tasks: [][]int{{1, 1}, {2, 1}, {3, 1}}, want: []int{0, 1, 2}},
		{tasks: [][]int{{4, 5}, {5, 1}}, want: []int{0, 1}},
		{tasks: [][]int{{4, 10}, {5, 4}, {6, 2}}, want: []int{0, 2, 1}},
		// fails
		{tasks: [][]int{{19, 13}, {16, 9}, {21, 10}, {32, 25}, {37, 4}, {49, 24}, {2, 15}, {38, 41}, {37, 34}, {33, 6}, {45, 4}, {18, 18}, {46, 39}, {12, 24}}, want: []int{6, 1, 2, 9, 4, 10, 0, 11, 5, 13, 3, 8, 12, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getOrder(tt.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
