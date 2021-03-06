package kdtree

import (
	"reflect"
	"testing"
)

func TestKDTree_Insert(t1 *testing.T) {
	tests := []struct {
		name   string
		points [][]int
		want   [][]int
	}{
		{
			"1",
			[][]int{{3, 6}, {17, 15}, {13, 15}, {6, 12},
				{9, 1}, {2, 7}, {10, 19}},
			[][]int{{2, 7}, {3, 6}, {6, 12}, {9, 1}, {17, 15}, {10, 19}, {13, 15}},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := &KDTree{K: 2}
			for i := range tt.points {
				t.Insert(tt.points[i])
			}
			var arr [][]int
			f := func(n *Node) {
				arr = append(arr, n.Point)
			}

			InOrderFunc(t.Root, f)

			if !reflect.DeepEqual(arr, tt.want) {
				t1.Errorf("got = %v, want %v", arr, tt.want)
			}
		})
	}
}
