package _1712_ways_to_split_array_into_three_subarrays

import (
	"fmt"
	"testing"
)

func Test_waysToSplit(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		//{[]int{1, 1, 1}, 1},
		//{[]int{1, 2, 2, 2, 5, 0}, 3},
		//{[]int{1, 2, 2, 2, 5, 0, 0}, 3},
		//{[]int{3, 2, 1}, 0},
		//{[]int{2, 5, 0}, 0},
		//{[]int{0, 0, 0}, 1},
		//{[]int{0, 0, 0, 0, 0, 0}, 10},
		{[]int{8892, 2631, 7212, 1188, 6580, 1690, 5950, 7425, 8787, 4361, 9849, 4063, 9496, 9140, 9986, 1058, 2734, 6961, 8855, 2567, 7683, 4770, 40, 850, 72, 2285, 9328, 6794, 8632, 9163, 3928, 6962, 6545, 6920, 926, 8885, 1570, 4454, 6876, 7447, 8264, 3123, 2980, 7276, 470, 8736, 3153, 3924, 3129, 7136, 1739, 1354, 661, 1309, 6231, 9890, 58, 4623, 3555, 3100, 3437}, 227},
		{bigArray, 290940717},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := waysToSplit(tt.nums); got != tt.want {
				t.Errorf("waysToSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}
