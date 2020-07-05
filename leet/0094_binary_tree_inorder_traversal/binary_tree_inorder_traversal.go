package _094_binary_tree_inorder_traversal

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"
)

var res []int

// TODO 1 - iterative
func inorderTraversal(root *btree.Node) []int {
	queue, res := make([]*btree.Node, 0), make([]int, 0)
	if root != nil {
		queue = append(queue, root)
	}
	var left bool
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		for node != nil {
			if node.Left != nil {
				node = node.Left
				left = true
			} else if node.Right != nil {
				node = node.Right
				left = false
			} else {
				node = nil
			}
			if node != nil {
				queue = append(queue, node)
				continue
			}
			last := len(queue) - 1
			res = append(res, queue[last].Val)
			queue = queue[:last]
			if len(queue) == 0 {
				return res
			}
			if left {
				queue[len(queue)-1].Left = nil
			} else {
				queue[len(queue)-1].Right = nil
			}
		}

	}
	return res
}

func inorderTraversalRecursive(root *btree.Node, res *[]int) {
	if root != nil {
		inorderTraversalRecursive(root.Left, res)
		*res = append(*res, root.Val)
		inorderTraversalRecursive(root.Right, res)
	}
}
