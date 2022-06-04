package _041_first_missing_positive

import (
	"fmt"
	"testing"
)

func Test_firstMissingPositive(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		//{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := firstMissingPositive(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
