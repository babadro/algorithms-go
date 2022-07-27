package _6_reconstructing_a_sequence

import (
	"fmt"
	"testing"
)

func Test_canConstruct(t *testing.T) {
	tests := []struct {
		originalSeq []int
		sequences   [][]int
		want        bool
	}{
		{[]int{1, 2, 3, 4}, [][]int{{1, 2}, {2, 3}, {3, 4}}, true},
		{[]int{1, 2, 3, 4}, [][]int{{1, 2}, {2, 3}, {2, 4}}, false},
		{[]int{3, 1, 4, 2, 5}, [][]int{{3, 1, 5}, {1, 4, 2, 5}}, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.originalSeq), func(t *testing.T) {
			if got := canConstruct(tt.originalSeq, tt.sequences); got != tt.want {
				t.Errorf("canConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
