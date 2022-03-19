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
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := recoverArray(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
