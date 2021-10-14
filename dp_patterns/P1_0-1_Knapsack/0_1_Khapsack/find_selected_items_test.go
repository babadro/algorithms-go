package __1_Khapsack

import (
	"reflect"
	"sort"
	"testing"
)

func TestFindSelectedElements(t *testing.T) {
	tests := []struct {
		dp       [][]int
		weights  []int
		profits  []int
		capacity int
		want     []int
	}{
		{dp: [][]int{
			{0, 1, 1, 1, 1, 1, 1, 1},
			{0, 1, 6, 7, 7, 7, 7, 7},
			{0, 1, 6, 10, 11, 16, 17, 17},
			{0, 1, 6, 10, 11, 16, 17, 22},
		},
			weights:  []int{1, 2, 3, 5},
			profits:  []int{1, 6, 10, 16},
			capacity: 7,
			want:     []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := FindSelectedElements(tt.dp, tt.weights, tt.profits, tt.capacity)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindSelectedElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
