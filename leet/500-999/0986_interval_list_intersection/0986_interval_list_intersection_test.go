package _986_interval_list_intersection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intervalIntersection(t *testing.T) {
	tests := []struct {
		firstList  [][]int
		secondList [][]int
		want       [][]int
	}{
		{[][]int{{0, 2}, {5, 10}, {13, 23}, {24, 25}},
			[][]int{{1, 5}, {8, 12}, {15, 24}, {25, 26}},
			[][]int{{1, 2}, {5, 5}, {8, 10}, {15, 23}, {24, 24}, {25, 25}},
		},
		{
			[][]int{{1, 3}, {5, 9}}, [][]int{}, nil,
		},
		{
			[][]int{{1, 2}, {4, 5}}, [][]int{{2, 3}}, [][]int{{2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v_%v", tt.firstList, tt.secondList), func(t *testing.T) {
			require.Equal(t, tt.want, intervalIntersection2(tt.firstList, tt.secondList))
		})
	}
}
