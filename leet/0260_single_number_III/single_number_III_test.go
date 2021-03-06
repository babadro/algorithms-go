package _260_single_number_III

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func Test_singleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{[]int{1, 0}, []int{1, 0}},
		{[]int{1, 0, 2, 2, 3, 3}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := singleNumber(tt.nums)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("singleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
