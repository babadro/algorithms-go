package __single_number

import (
	"fmt"
	"testing"
)

func Test_findSingleNumber(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 4, 2, 1, 3, 2, 3}, 4},
		{[]int{7, 9, 7}, 9},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			if got := findSingleNumber(tt.arr); got != tt.want {
				t.Errorf("findSingleNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
