package __find_all_duplicate_numbers

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_findAllDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{3, 4, 4, 5, 5}, []int{4, 5}},
		{[]int{5, 4, 7, 2, 3, 5, 3}, []int{3, 5}},
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
		{[]int{1, 1, 2}, []int{1}},
		{[]int{1}, nil},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := findAllDuplicates(tt.nums)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
