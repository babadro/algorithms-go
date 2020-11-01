package _897_increasing_order_search_tree

import "github.com/babadro/algorithms-go/04_TreesAndGraphs/btree"

// tptl. Simple but not in place (O(N))
func increasingBST(root *btree.Node) *btree.Node {
	var arr []*btree.Node

	inOrder(root, &arr)

	if len(arr) == 0 {
		return nil
	}

	node, res := arr[0], arr[0]
	node.Left, node.Right = nil, nil
	for _, n := range arr[1:] {
		n.Left, n.Right = nil, nil
		node.Right = n
		node = n
	}

	return res
}

func inOrder(root *btree.Node, arr *[]*btree.Node) {
	if root != nil {
		inOrder(root.Left, arr)
		*arr = append(*arr, root)
		inOrder(root.Right, arr)
	}
}
