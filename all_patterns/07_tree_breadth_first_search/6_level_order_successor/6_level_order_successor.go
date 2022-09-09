package __level_order_successor

import "github.com/babadro/algorithms-go/base/binaryTree"

func findSuccessor(root *binaryTree.Node, key int) *binaryTree.Node {
	if root == nil {
		return nil
	}

	q := []*binaryTree.Node{root}
	var node *binaryTree.Node
	for len(q) > 0 {
		node, q = q[0], q[1:]
		if node.Left != nil {
			q = append(q, node.Left)
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}

		if node.Val == key {
			break
		}
	}

	if len(q) != 0 {
		return q[0]
	}

	return nil
}
