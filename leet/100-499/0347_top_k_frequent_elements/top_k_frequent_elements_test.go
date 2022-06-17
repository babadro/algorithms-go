package _347_top_k_frequent_elements

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_topKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{3, 0, 1, 0}, 1, []int{0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := topKFrequent2(tt.nums, tt.k)
			sort.Ints(tt.want)
			sort.Ints(got)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("topKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
