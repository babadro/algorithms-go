package _095_unique_binary_search_trees_ii

import "github.com/babadro/algorithms-go/base/binaryTree"

// dyx passed. medium (hard for me)
func generateTrees(n int) []*binaryTree.Node {
	return rec(1, n)
}

func rec(start, end int) []*binaryTree.Node {
	if start > end {
		return []*binaryTree.Node{nil}
	}

	var res []*binaryTree.Node
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
