package _2197_replace_non_coprime_numbers_in_array

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_replaceNonCoprimes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{6, 4, 3, 2, 7, 6, 2}, []int{12, 7, 6}},
		{[]int{2, 2, 1, 1, 3, 3, 3}, []int{2, 1, 1, 3}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := replaceNonCoprimes(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceNonCoprimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
