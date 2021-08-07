package _1932_merge_bsts_to_create_single_bst

import (
	"github.com/babadro/algorithms-go/base/binaryTree"
	"math"
)

// doesn't work
func canMerge2(trees []*binaryTree.Node) *binaryTree.Node {
	roots := make(map[int]*binaryTree.Node)
	leafToParent := make(map[int]*binaryTree.Node)

	// fill maps
	for _, tree := range trees {
		roots[tree.Val] = tree

		if tree.Left == nil && tree.Right == nil {
			continue
		}

		if tree.Left != nil {
			if _, ok := leafToParent[tree.Left.Val]; ok {
				return nil
			}

			leafToParent[tree.Left.Val] = tree
		}

		if tree.Right != nil {
			if _, ok := leafToParent[tree.Right.Val]; ok {
				return nil
			}

			leafToParent[tree.Right.Val] = tree
		}
	}

	var parent *binaryTree.Node
	var ok bool
	for len(roots) > 1 {
		rootK, rootV := pop(roots)
		if parent, ok = leafToParent[rootK]; !ok {
			return nil
		}

		if parent.Left != nil && parent.Left.Val == rootK {
			parent.Left = rootV
		} else {
			parent.Right = rootV
		}

		if rootV.Left != nil || rootV.Right != nil {
			delete(leafToParent, rootK)
		}
	}

	for _, root := range roots {
		if !isValidBST(root) {
			return nil
		}

		return root
	}

	return nil
}

func pop(nodes map[int]*binaryTree.Node) (k int, v *binaryTree.Node) {
	for k, v = range nodes {
		delete(nodes, k)

		return k, v
	}

	return 0, nil
}

func isValidBST(root *binaryTree.Node) bool {
	return recursive(root, math.MinInt64, math.MaxInt64)
}

func recursive(node *binaryTree.Node, min, max int) bool {
	if node == nil {
		return true
	}
	if node.Val <= min || node.Val >= max {
		return false
	}

	return recursive(node.Left, min, node.Val) && recursive(node.Right, node.Val, max)
}
