package _897_increasing_order_search_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

var curr *btree.Node

// passed. in place. tptl
func increasingBST2(root *btree.Node) *btree.Node {
	res := &btree.Node{}
	curr = res
	inOrder2(root)
	return res.Right
}

func inOrder2(root *btree.Node) {
	if root != nil {
		inOrder2(root.Left)
		curr.Right = root
		curr = curr.Right
		root.Left = nil
		inOrder2(root.Right)
	}
}
