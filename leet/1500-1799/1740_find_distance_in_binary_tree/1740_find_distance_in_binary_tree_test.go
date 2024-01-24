package _1740_find_distance_in_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"testing"
)

func Test_findDistance(t *testing.T) {
	tests := []struct {
		builder func() *binaryTree.Node
		p       int
		q       int
		want    int
	}{
		{
			builder: func() *binaryTree.Node {
				n3, n5, n1 := n(3), n(5), n(1)
				n6, n2, n0, n8 := n(6), n(2), n(0), n(8)
				n7, n4 := n(7), n(4)

				n3.Left, n3.Right = n5, n1
				n5.Left, n5.Right = n6, n2
				n1.Left, n1.Right = n0, n8
				n2.Left, n2.Right = n7, n4

				return n3
			},
			p: 5, q: 0, want: 3,
		},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findDistance(tt.builder(), tt.p, tt.q); got != tt.want {
				t.Errorf("findDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(v int) *binaryTree.Node {
	return &binaryTree.Node{Val: v}
}
