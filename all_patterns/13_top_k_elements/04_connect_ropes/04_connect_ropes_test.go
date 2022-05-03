package _4_connect_ropes

import (
	"fmt"
	"testing"
)

func Test_minimumCostToConnectRopes(t *testing.T) {
	tests := []struct {
		ropeLengths []int
		want        int
	}{
		{[]int{3, 4, 5, 6}, 36},
		{[]int{1, 3, 11, 5, 2}, 42},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.ropeLengths), func(t *testing.T) {
			if got := minimumCostToConnectRopes(tt.ropeLengths); got != tt.want {
				t.Errorf("minimumCostToConnectRopes() = %v, want %v", got, tt.want)
			}
		})
	}
}
