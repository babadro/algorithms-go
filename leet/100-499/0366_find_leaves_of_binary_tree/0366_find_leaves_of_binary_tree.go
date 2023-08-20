package _366_find_leaves_of_binary_tree

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"github.com/babadro/algorithms-go/utils"
)

// bnsrg passed
func findLeaves(root *binaryTree.Node) [][]int {
	res := [][]int{nil}

	postOrder(root, &res)

	return res[1:]
}

func postOrder(node *binaryTree.Node, arr *[][]int) int {
	if node == nil {
		return 0
	}

	leftDepth, rightDepth := postOrder(node.Left, arr), postOrder(node.Right, arr)

	depth := utils.Max(leftDepth, rightDepth) + 1

	if depth == len(*arr) {
		*arr = append(*arr, []int{node.Val})
	} else {
		(*arr)[depth] = append((*arr)[depth], node.Val)
	}

	return depth
}
