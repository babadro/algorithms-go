package _834_sum_of_distances_in_tree

import (
	"reflect"
	"testing"
)

func Test_sumOfDistancesInTree(t *testing.T) {
	tests := []struct {
		n     int
		edges [][]int
		want  []int
	}{
		{
			6, [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}}, []int{8, 12, 6, 10, 10, 10},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := sumOfDistancesInTree(tt.n, tt.edges); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumOfDistancesInTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
