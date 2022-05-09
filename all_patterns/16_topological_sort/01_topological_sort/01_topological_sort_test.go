package _1_topological_sort

import (
	"reflect"
	"testing"
)

func Test_topologicalSort(t *testing.T) {
	tests := []struct {
		vertices int
		edges    [][]int
		want     []int
	}{
		{4, [][]int{{3, 2}, {3, 0}, {2, 0}, {2, 1}}, []int{3, 2, 0, 1}},
		{5, [][]int{{4, 2}, {4, 3}, {2, 0}, {2, 1}, {3, 1}}, []int{4, 2, 3, 0, 1}},
		{7, [][]int{{6, 4}, {6, 2}, {5, 3}, {5, 4}, {3, 0}, {3, 1}, {3, 2}, {4, 1}}, []int{5, 6, 3, 4, 0, 2, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := topologicalSort(tt.vertices, tt.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topologicalSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
