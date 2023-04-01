package _111_minimum_depth_of_binary_tree

import (
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
	"github.com/stretchr/testify/require"
)

func Test_minDepth(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want int
	}{
		{"1", []int{3, 9, 20, binaryTree.Null, binaryTree.Null, 15, 7}, 2},
		//{"2", []int{2, 3, binaryTree.Null, 4, binaryTree.Null, binaryTree.Null, binaryTree.Null, 5, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := binaryTree.ArrayToBinaryTree(tt.arr)
			if got := minDepth(root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepth2(t *testing.T) {
	n2 := &binaryTree.Node{Val: 2}
	n3 := &binaryTree.Node{Val: 3}
	n4 := &binaryTree.Node{Val: 4}
	n5 := &binaryTree.Node{Val: 5}
	n6 := &binaryTree.Node{Val: 6}

	n2.Left = n3
	n3.Left = n4
	n4.Left = n5
	n5.Left = n6

	t.Log(minDepth(n2))
}

func Test3(t *testing.T) {
	n1, n2, n3, n4, n5 := n(1), n(2), n(3), n(4), n(5)
	n1.Left, n1.Right = n2, n3

	n3.Left, n3.Right = n4, n5

	require.Equal(t, 2, minDepth4(n1))
}

func Test4(t *testing.T) {
	n1, n2, n3, n4, n5 := n(1), n(2), n(3), n(4), n(5)
	n1.Left, n1.Right = n2, n3

	n2.Left, n2.Right = n4, n5

	require.Equal(t, 2, minDepth4(n1))
}

func Test4_2(t *testing.T) {
	n2, n3, n4, n5, n6 := n(2), n(3), n(4), n(5), n(6)

	n2.Right = n3
	n3.Right = n4
	n4.Right = n5
	n5.Right = n6

	require.Equal(t, 5, minDepth4(n2))
}

func n(val int) *binaryTree.Node {
	return &binaryTree.Node{Val: val}
}
