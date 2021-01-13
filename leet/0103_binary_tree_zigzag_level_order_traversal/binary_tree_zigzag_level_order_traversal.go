package _103_binary_tree_zigzag_level_order_traversal

import "github.com/babadro/algorithms-go/base/binaryTree"

// 100% 72%
func zigzagLevelOrder(root *binaryTree.Node) [][]int {
	var res [][]int

	addValueToLevel(root, 0, &res)

	for i := range res {
		if i%2 != 0 {
			n := len(res[i])
			for j := 0; j < n/2; j++ {
				res[i][j], res[i][n-1-j] = res[i][n-1-j], res[i][j]
			}
		}
	}

	return res
}

func addValueToLevel(node *binaryTree.Node, level int, levels *[][]int) {
	if node == nil {
		return
	}
	if level == len(*levels) {
		*levels = append(*levels, []int{node.Val})
	} else {
		(*levels)[level] = append((*levels)[level], node.Val)
	}
	addValueToLevel(node.Left, level+1, levels)
	addValueToLevel(node.Right, level+1, levels)
}

// Clear, single pass, but waste of memory. 100% 27%
func zigzagLevelOrder2(root *binaryTree.Node) [][]int {
	var res [][]int

	addValueToLevel(root, 0, &res)

	return res
}

func addValueToLevel2(node *binaryTree.Node, level int, levels *[][]int) {
	if node == nil {
		return
	}
	if level == len(*levels) {
		*levels = append(*levels, []int{node.Val})
	} else {
		if level%2 == 0 {
			(*levels)[level] = append((*levels)[level], node.Val)
		} else {
			(*levels)[level] = append([]int{node.Val}, (*levels)[level]...)
		}

	}
	addValueToLevel2(node.Left, level+1, levels)
	addValueToLevel2(node.Right, level+1, levels)
}
