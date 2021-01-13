package _094_binary_tree_inorder_traversal

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
)

// TODO 2 - need to understand. Not mine
func inorderTraversal(root *binaryTree.Node) []int {
	if root == nil {
		return nil
	}
	stack, res := make([]*binaryTree.Node, 0), make([]int, 0)
	node := root
	for node != nil || len(stack) > 0 { // TODO 2: interesting. Need to understand
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		last := len(stack) - 1
		node = stack[last]
		stack = stack[:last]
		res = append(res, node.Val)
		node = node.Right
	}
	return res
}

// Doesn't work
func inorderTraversal2(root *binaryTree.Node) []int {
	queue, res := make([]*binaryTree.Node, 0), make([]int, 0)
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

func inorderTraversalRecursive(root *binaryTree.Node, res *[]int) {
	if root != nil {
		inorderTraversalRecursive(root.Left, res)
		*res = append(*res, root.Val)
		inorderTraversalRecursive(root.Right, res)
	}
}
