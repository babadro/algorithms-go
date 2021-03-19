package _645_set_mismatch

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findErrorNums(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 2, 4}, []int{2, 3}},
		{[]int{1, 1}, []int{1, 2}},
		{[]int{2, 2, 3, 4}, []int{2, 1}},
		{[]int{3, 2, 3, 4, 6, 5}, []int{3, 1}},
		{[]int{8, 7, 3, 5, 3, 6, 1, 4}, []int{3, 2}},
		{[]int{1, 5, 3, 2, 2, 7, 6, 4, 8, 9}, []int{2, 10}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findErrorNums3(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findErrorNums() = %v, want %v", got, tt.want)
			}
		})
	}
}
