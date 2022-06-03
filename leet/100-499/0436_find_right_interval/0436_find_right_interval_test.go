package _436_find_right_interval

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findRightInterval(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      []int
	}{
		{[][]int{{1, 2}}, []int{-1}},
		{[][]int{{3, 4}, {2, 3}, {1, 2}}, []int{-1, 0, 1}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.intervals), func(t *testing.T) {
			if got := findRightInterval(tt.intervals); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findRightInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
