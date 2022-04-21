package _0_introduction

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_avgOfSubarrayOfSizeK(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want []float64
	}{
		{[]int{1, 3, 2, 6, -1, 4, 1, 8, 2}, 5, []float64{2.2, 2.8, 2.4, 3.6, 2.8}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d, %v", tt.k, tt.arr), func(t *testing.T) {
			if got := avgOfSubarrayOfSizeK(tt.k, tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("avgOfSubarrayOfSizeK() = %v, want %v", got, tt.want)
			}
		})
	}
}
