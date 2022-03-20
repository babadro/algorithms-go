package _2122_recover_the_original_array

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_recoverArray(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{2, 10, 6, 4, 8, 12}, []int{3, 7, 11}},
		{[]int{1, 1, 3, 3}, []int{2, 2}},
		{[]int{5, 435}, []int{220}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := recoverArray3(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
