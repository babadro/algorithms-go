package _102_binary_tree_level_order_traversal

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/binaryTree"

// 100% 66%
// TODO 2 find better solution in discuss
func levelOrder(root *binaryTree.Node) [][]int {
	if root == nil {
		return nil
	}
	nodeArrays := [][]*binaryTree.Node{{root}}
	for {
		last := len(nodeArrays) - 1
		var arr []*binaryTree.Node
		for _, node := range nodeArrays[last] {
			if node.Left != nil {
				arr = append(arr, node.Left)
			}
			if node.Right != nil {
				arr = append(arr, node.Right)
			}
		}
		if arr == nil {
			break
		}
		nodeArrays = append(nodeArrays, arr)
	}

	res := make([][]int, len(nodeArrays))
	for i := range nodeArrays {
		res[i] = make([]int, len(nodeArrays[i]))
		for j := range nodeArrays[i] {
			res[i][j] = nodeArrays[i][j].Val
		}
	}

	return res
}
