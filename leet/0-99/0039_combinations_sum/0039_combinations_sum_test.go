package _039_combinations_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_combinationSum(t *testing.T) {
	tests := []struct {
		candidates []int
		target     int
		want       [][]int
	}{
		{[]int{2, 3, 5}, 8, [][]int{{3, 5}, {2, 3, 3}, {2, 2, 2, 2}}},
		{[]int{2, 3, 5}, 8, [][]int{{3, 5}, {2, 3, 3}, {2, 2, 2, 2}}},
		{[]int{2}, 1, nil},
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1}, 2, [][]int{{1, 1}}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%d", tt.candidates, tt.target), func(t *testing.T) {
			if got := combinationSum(tt.candidates, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
