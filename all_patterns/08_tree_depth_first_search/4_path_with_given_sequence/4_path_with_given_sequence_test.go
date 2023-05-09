package __path_with_given_sequence

import (
	"fmt"
	"testing"

	"github.com/babadro/algorithms-go/base/binaryTree"
)

func Test_findPath(t *testing.T) {
	type rootWrapper struct {
		root *binaryTree.Node
	}

	n1_1, n2_7, n2_9, n3_2, n3_9 := n(1), n(7), n(9), n(2), n(9)
	n1_1.Left, n1_1.Right = n2_7, n2_9
	n2_9.Left, n2_9.Right = n3_2, n3_9
	root1 := rootWrapper{root: n1_1}

	n1_1, n2_0, n2_1, n3_1, n3_6, n3_5 := n(1), n(0), n(1), n(1), n(6), n(5)
	n1_1.Left, n1_1.Right = n2_0, n2_1
	n2_0.Left = n3_1
	n2_1.Left, n2_1.Right = n3_6, n3_5
	root2 := rootWrapper{root: n1_1}

	n1_1, n2_7, n2_9, n3_9 = n(1), n(7), n(9), n(9)
	n1_1.Left, n1_1.Right = n2_7, n2_9
	n2_9.Right = n3_9
	root3 := rootWrapper{root: n1_1}

	tests := []struct {
		root rootWrapper
		seq  []int
		want bool
	}{
		{root1, []int{1, 9, 9}, true},
		{root1, []int{1, 9, 9, 5}, false},
		{root2, []int{1, 1, 6}, true},
		{root3, []int{1, 9}, false},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.seq), func(t *testing.T) {
			if got := findPath(tt.root.root, tt.seq); got != tt.want {
				t.Errorf("findPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func n(v int) *binaryTree.Node {
	return &binaryTree.Node{Val: v}
}
