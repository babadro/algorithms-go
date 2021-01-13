package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"testing"
)

func Test_buildTree(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}

	root := buildTree2(preorder, inorder)

	var res []int
	binaryTree.InOrder(root, &res)

	t.Log(res)
}
