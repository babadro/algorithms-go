package _1_pair_with_target_sum

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_search(t *testing.T) {
	tests := []struct {
		sortedArr []int
		target    int
		want      []int
	}{
		{[]int{1, 2, 3, 4, 6}, 6, []int{1, 3}},
		{[]int{2, 5, 9, 11}, 11, []int{0, 2}},
		{[]int{1, 2, 3, 4, 5}, 10, nil},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v, %d", tt.sortedArr, tt.target), func(t *testing.T) {
			if got := searchHashMap(tt.sortedArr, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
