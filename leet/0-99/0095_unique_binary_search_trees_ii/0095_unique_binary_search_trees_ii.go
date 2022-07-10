package _095_unique_binary_search_trees_ii

import "github.com/babadro/algorithms-go/base/binaryTree"

// dyx passed. medium (hard for me)
func generateTrees(n int) []*binaryTree.Node {
	return rec(1, n)
}

func rec(start, end int) []*binaryTree.Node {
	var res []*binaryTree.Node
	if start > end {
		res = append(res, nil)
		return res
	}

	for i := start; i <= end; i++ {
		leftSubTrees := rec(start, i-1)
		rightSubtrees := rec(i+1, end)
		for _, leftTree := range leftSubTrees {
			for _, rightTree := range rightSubtrees {
				root := &binaryTree.Node{Val: i}
				root.Left, root.Right = leftTree, rightTree

				res = append(res, root)
			}
		}
	}

	return res
}
