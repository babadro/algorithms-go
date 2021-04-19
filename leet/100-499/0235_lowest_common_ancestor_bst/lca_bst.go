package _235_lowest_common_ancestor_bst

import (
	"github.com/babadro/algorithms-go/leet/100-499/0108_sorted_array_to_BST"
)

// recursive
func lowestCommonAncestor(root, p, q *_0108_sorted_array_to_BST.TreeNode) *_0108_sorted_array_to_BST.TreeNode {
	if root == nil {
		return nil
	}
	parentVal := root.Val
	pVal := p.Val
	qVal := q.Val

	if pVal > parentVal && qVal > parentVal {
		return lowestCommonAncestor(root.Right, p, q)
	} else if pVal < parentVal && qVal < parentVal {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return root
	}
}

// iterative
func lowestCommonAncestorIterative(root, p, q *_0108_sorted_array_to_BST.TreeNode) *_0108_sorted_array_to_BST.TreeNode {
	pVal := p.Val
	qVal := q.Val
	node := root

	for node != nil {
		nodeVal := node.Val
		if pVal > nodeVal && qVal > nodeVal {
			return lowestCommonAncestorIterative(node.Right, p, q)
		} else if pVal < nodeVal && qVal < nodeVal {
			return lowestCommonAncestorIterative(node.Left, p, q)
		} else {
			return node
		}
	}
	return node
}
