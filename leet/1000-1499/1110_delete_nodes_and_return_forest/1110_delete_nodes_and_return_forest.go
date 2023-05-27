package _1110_delete_nodes_and_return_forest

import "github.com/babadro/algorithms-go/base/binaryTree"

// #bnsrg
// todo 2 find shorter soluiton?
func delNodes(root *binaryTree.Node, to_delete []int) []*binaryTree.Node {
	toDelete := make(map[int]bool, len(to_delete))
	for _, val := range to_delete {
		toDelete[val] = true
	}

	valToPtr := make(map[int]*binaryTree.Node)
	rec(root, valToPtr)

	resM := map[int]bool{root.Val: true}

	for val := range toDelete {
		if resM[val] {
			delete(resM, val)
		}

		node := valToPtr[val]

		if node.Left != nil && !toDelete[node.Left.Val] {
			resM[node.Left.Val] = true
		}

		if node.Right != nil && !toDelete[node.Right.Val] {
			resM[node.Right.Val] = true
		}
	}

	for _, node := range valToPtr {
		if node.Left != nil && toDelete[node.Left.Val] {
			node.Left = nil
		}

		if node.Right != nil && toDelete[node.Right.Val] {
			node.Right = nil
		}
	}

	res := make([]*binaryTree.Node, 0, len(resM))
	for val := range resM {
		res = append(res, valToPtr[val])
	}

	return res
}

func rec(node *binaryTree.Node, m map[int]*binaryTree.Node) {
	if node == nil {
		return
	}

	m[node.Val] = node

	rec(node.Left, m)
	rec(node.Right, m)
}
