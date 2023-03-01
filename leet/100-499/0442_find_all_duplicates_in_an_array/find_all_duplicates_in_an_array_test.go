package _442_find_all_duplicates_in_an_array

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_findDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{4, 3, 2, 7, 8, 2, 3, 1}, []int{2, 3}},
		{[]int{1, 1, 2}, []int{1}},
		{[]int{1}, nil},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.nums), func(t *testing.T) {
			got := findDuplicates2(tt.nums)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
