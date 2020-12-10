package _110_balanced_binary_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

// passed. tptl. best solution
func isBalanced2(root *btree.Node) bool {
	return postOrder2(root) != -1
}

func postOrder2(root *btree.Node) int {
	if root == nil {
		return 0
	}

	h1 := postOrder2(root.Left)
	if h1 == -1 {
		return -1
	}

	h2 := postOrder2(root.Right)
	if h2 == -1 {
		return -1
	}

	if h1 < h2 {
		h1, h2 = h2, h1
	}

	if h1-h2 > 1 {
		return -1
	}

	return h1 + 1
}

// passed.
func isBalanced(root *btree.Node) bool {
	balanced := true
	postOrder(root, &balanced)

	return balanced
}

func postOrder(root *btree.Node, balanced *bool) int {
	if !*balanced || root == nil {
		return 0
	}

	h1 := postOrder(root.Left, balanced)
	h2 := postOrder(root.Right, balanced)
	if h1 < h2 {
		h2, h1 = h1, h2
	}
	diff := h1 - h2
	if diff > 1 {
		*balanced = false
	}

	return h1 + 1
}
