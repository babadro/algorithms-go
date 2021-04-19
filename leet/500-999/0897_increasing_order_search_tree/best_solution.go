package _897_increasing_order_search_tree

import "github.com/babadro/algorithms-go/base/binaryTree"

var curr *binaryTree.Node

// passed. in place. tptl
func increasingBST2(root *binaryTree.Node) *binaryTree.Node {
	res := &binaryTree.Node{}
	curr = res
	inOrder2(root)
	return res.Right
}

func inOrder2(root *binaryTree.Node) {
	if root != nil {
		inOrder2(root.Left)
		curr.Right = root
		curr = curr.Right
		root.Left = nil
		inOrder2(root.Right)
	}
}
