package kdtree

import (
	"fmt"
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
				t1.Errorf("got = %v, found %v", arr, tt.want)
			}
		})
	}
}

func TestKDTree_Search(t1 *testing.T) {
	tests := []struct {
		point       []int
		treeBuilder func() *KDTree
		found       bool
	}{
		{
			point: []int{10, 19},
			treeBuilder: func() *KDTree {
				t := &KDTree{K: 2}
				for _, point := range [][]int{{3, 6}, {17, 15}, {13, 15}, {6, 12}, {9, 1}, {2, 7}, {10, 19}} {
					t.Insert(point)
				}

				return t
			},
			found: true,
		},
		{
			point: []int{12, 19},
			treeBuilder: func() *KDTree {
				t := &KDTree{K: 2}
				for _, point := range [][]int{{3, 6}, {17, 15}, {13, 15}, {6, 12}, {9, 1}, {2, 7}, {10, 19}} {
					t.Insert(point)
				}

				return t
			},
			found: false,
		},
	}
	for _, tt := range tests {
		t1.Run(fmt.Sprintf("%v", tt.point), func(t1 *testing.T) {
			t := tt.treeBuilder()
			got := t.Search(tt.point)
			if !tt.found && got != nil {
				t1.Errorf("Search() = %v, want nil", got)
			} else if tt.found && got == nil {
				t1.Errorf("Search() = nil, want %v", tt.point)
			} else if tt.found && got != nil && !reflect.DeepEqual(tt.point, got.Point) {
				t1.Errorf("got = %v, want %v", got, tt.point)
			}
		})
	}
}

// fails
func TestKDTree_Delete(t1 *testing.T) {
	tests := []struct {
		pointToDelete []int
		treeBuilder   func() *KDTree
		want          [][]int
	}{
		{
			pointToDelete: []int{30, 40},
			treeBuilder: func() *KDTree {
				t := &KDTree{K: 2}
				for _, pointToDelete := range [][]int{{30, 40}, {5, 25}, {70, 70}, {10, 12}, {50, 30}, {35, 45}} {
					t.Insert(pointToDelete)
				}

				return t
			},
			want: [][]int{{10, 12}, {5, 25}, {35, 45}, {50, 30}, {70, 70}},
		},
		{
			pointToDelete: []int{70, 70},
			treeBuilder: func() *KDTree {
				t := &KDTree{K: 2}
				for _, point := range [][]int{{30, 40}, {5, 25}, {70, 70}, {10, 12}, {50, 30}, {35, 45}} {
					t.Insert(point)
				}

				return t
			},
			want: [][]int{{10, 12}, {5, 25}, {30, 40}, {50, 30}, {35, 45}},
		},
	}
	for _, tt := range tests {
		t1.Run(fmt.Sprintf("%v", tt.pointToDelete), func(t1 *testing.T) {
			t := tt.treeBuilder()
			t.DeleteNode(tt.pointToDelete)

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
