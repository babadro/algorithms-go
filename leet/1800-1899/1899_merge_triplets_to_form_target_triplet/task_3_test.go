package _899_merge_triplets_to_form_target_triplet

import (
	"fmt"
	"testing"
)

func Test_mergeTriplets(t *testing.T) {
	tests := []struct {
		triplets [][]int
		target   []int
		want     bool
	}{
		{[][]int{{2, 5, 3}, {1, 8, 4}, {1, 7, 5}}, []int{2, 7, 5}, true},
		{[][]int{{1, 3, 4}, {2, 5, 8}}, []int{2, 5, 8}, true},
		{[][]int{{2, 5, 3}, {2, 3, 4}, {1, 2, 5}, {5, 2, 3}}, []int{5, 5, 5}, true},
		{[][]int{{3, 4, 5}, {4, 5, 6}}, []int{3, 2, 5}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v %d", tt.triplets, tt.target), func(t *testing.T) {
			if got := mergeTriplets(tt.triplets, tt.target); got != tt.want {
				t.Errorf("mergeTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
