package _013_partition_array_into_three_parts_with_equal_sum

import (
	"fmt"
	"testing"
)

func Test_canThreePartsEqualSum(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{0, 2, 1, -6, 6, -7, 9, 1, 2, 0, 1}, true},
		{[]int{0, 2, 1, -6, 6, 7, 9, -1, 2, 0, 1}, false},
		{[]int{3, 3, 6, 5, -2, 2, 5, 1, -9, 4}, true},
		{[]int{1, 2, 3, 3}, true},
		{[]int{1, -1, 1, -1}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.arr), func(t *testing.T) {
			if got := canThreePartsEqualSum2(tt.arr); got != tt.want {
				t.Errorf("canThreePartsEqualSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
