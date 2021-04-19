package _1835_find_xor_sum_of_all_pairs_bitwise_and

import "testing"

func Test_getXORSum(t *testing.T) {
	tests := []struct {
		name string
		arr1 []int
		arr2 []int
		want int
	}{
		{arr1: []int{1, 2, 3}, arr2: []int{6, 5}, want: 0},
		{arr1: []int{1}, arr2: []int{3}, want: 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getXORSum(tt.arr1, tt.arr2); got != tt.want {
				t.Errorf("getXORSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
