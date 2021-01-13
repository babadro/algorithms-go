package _105_construct_binary_tree_from_preorder_and_inorder_traversal

import (
	"github.com/babadro/algorithms-go/04_TreesAndGraphs/binaryTree"
)

// best recursive solution.
// todo 3 look at iterative solution
func recursiveHelper(preorder []int, preorderStart, preorderEnd, inorderStart int, m map[int]int) *binaryTree.Node {
	if preorderStart > preorderEnd {
		return nil
	}

	rootIdx := m[preorder[preorderStart]]
	leftLen := rootIdx - inorderStart
	root := &binaryTree.Node{Val: preorder[preorderStart]}

	root.Left = recursiveHelper(preorder, preorderStart+1, preorderStart+leftLen, inorderStart, m)
	root.Right = recursiveHelper(preorder, preorderStart+leftLen+1, preorderEnd, rootIdx+1, m)

	return root
}

func buildTree2(preorder []int, inorder []int) *binaryTree.Node {
	m := make(map[int]int, len(preorder))

	for i, v := range inorder {
		m[v] = i
	}

	return recursiveHelper(preorder, 0, len(preorder)-1, 0, m)
}

// 96% 74%
func buildTree(preorder []int, inorder []int) *binaryTree.Node {
	if len(preorder) == 0 {
		return nil
	}

	root := &binaryTree.Node{Val: preorder[0]}

	leftLen := 0
	for i, v := range inorder {
		if v == preorder[0] {
			leftLen = i
			break
		}
	}

	root.Left = buildTree(preorder[1:leftLen+1], inorder[:leftLen])
	root.Right = buildTree(preorder[leftLen+1:], inorder[leftLen+1:])

	return root
}
