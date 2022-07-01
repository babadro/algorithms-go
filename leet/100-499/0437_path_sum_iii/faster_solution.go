package _437_path_sum_iii

import "github.com/babadro/algorithms-go/base/binaryTree"

// dyx. passed. faster and shorter
func pathSum2(root *binaryTree.Node, targetSum int) int {
	var curPath []int
	return rec(root, targetSum, &curPath)
}

func rec(node *binaryTree.Node, targetSum int, curPath *[]int) int {
	if node == nil {
		return 0
	}

	*curPath = append(*curPath, node.Val)
	pathCount, sum := 0, 0
	for i := len(*curPath) - 1; i >= 0; i-- {
		sum += (*curPath)[i]
		if sum == targetSum {
			pathCount++
		}
	}

	pathCount += rec(node.Left, targetSum, curPath)
	pathCount += rec(node.Right, targetSum, curPath)

	*curPath = (*curPath)[:len(*curPath)-1]

	return pathCount
}
