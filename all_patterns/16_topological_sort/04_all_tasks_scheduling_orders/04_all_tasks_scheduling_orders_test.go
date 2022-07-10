package _04_all_tasks_scheduling_orders

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/babadro/algorithms-go/slices"
)

func Test_allOrders(t *testing.T) {
	tests := []struct {
		tasks         int
		prerequisites [][]int
		want          [][]int
	}{
		{3, [][]int{{0, 1}, {1, 2}}, [][]int{{0, 1, 2}}},
		{6, [][]int{{2, 5}, {0, 5}, {0, 4}, {1, 4}, {3, 2}, {1, 3}},
			[][]int{
				{0, 1, 4, 3, 2, 5},
				{0, 1, 3, 4, 2, 5},
				{0, 1, 3, 2, 4, 5},
				{0, 1, 3, 2, 5, 4},
				{1, 0, 3, 4, 2, 5},
				{1, 0, 3, 2, 4, 5},
				{1, 0, 3, 2, 5, 4},
				{1, 0, 4, 3, 2, 5},
				{1, 3, 0, 2, 4, 5},
				{1, 3, 0, 2, 5, 4},
				{1, 3, 0, 4, 2, 5},
				{1, 3, 2, 0, 5, 4},
				{1, 3, 2, 0, 4, 5},
			}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.prerequisites), func(t *testing.T) {
			got := allOrders(tt.tasks, tt.prerequisites)

			slices.Sort2DSlice(tt.want)
			slices.Sort2DSlice(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
