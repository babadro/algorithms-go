package _450_delete_node_in_a_bst

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

// passed. tptl
func deleteNode(root *btree.Node, key int) *btree.Node {
	if root == nil {
		return nil
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			return root.Right
		}
		if root.Right == nil {
			return root.Left
		}

		smallestRight := root.Right
		for smallestRight.Left != nil {
			smallestRight = smallestRight.Left
		}

		smallestRight.Left = root.Left

		return root.Right
	}

	return root
}
