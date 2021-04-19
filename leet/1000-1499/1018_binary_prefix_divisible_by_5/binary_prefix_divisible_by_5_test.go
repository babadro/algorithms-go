package _1018_binary_prefix_divisible_by_5

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_prefixesDivBy5(t *testing.T) {
	tests := []struct {
		A    []int
		want []bool
	}{
		{[]int{0, 1, 1}, []bool{true, false, false}},
		{[]int{1, 1, 1}, []bool{false, false, false}},
		{[]int{0, 1, 1, 1, 1, 1}, []bool{true, false, false, false, true, false}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.A), func(t *testing.T) {
			if got := prefixesDivBy5(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixesDivBy5() = %v, want %v", got, tt.want)
			}
		})
	}
}
