package _366_find_leaves_of_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"github.com/babadro/algorithms-go/utils"
)

// bnsrg passed
func findLeaves(root *binaryTree.Node) [][]int {
	arr := [][]*binaryTree.Node{nil}

	postOrder(root, &arr)

	var res [][]int
	for _, depth := range arr[1:] {
		values := make([]int, len(depth))

		for i, node := range depth {
			values[i] = node.Val
		}

		res = append(res, values)
	}

	return res
}

func postOrder(node *binaryTree.Node, arr *[][]*binaryTree.Node) int {
	if node == nil {
		return 0
	}

	leftDepth, rightDepth := postOrder(node.Left, arr), postOrder(node.Right, arr)

	depth := utils.Max(leftDepth, rightDepth) + 1

	if depth == len(*arr) {
		*arr = append(*arr, []*binaryTree.Node{node})
	} else {
		(*arr)[depth] = append((*arr)[depth], node)
	}

	return depth
}
