package _1325_delete_leaves_with_a_given_value

import "testing"

func TestRemoveLeafNode(t *testing.T) {
	//root := &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 2}
	//root.Right = &TreeNode{Val: 3}
	//root.Left.Left = &TreeNode{Val: 2}
	//root.Right.Left = &TreeNode{Val: 2}
	//root.Right.Right = &TreeNode{Val: 4}
	//
	//InOrderTraversal(t, root)
	//
	//result := removeLeafNodes(root, 2)
	//
	//InOrderTraversal(t, result)

	//root := &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 1}
	//root.Right = &TreeNode{Val: 1}
	//
	//InOrderTraversal(t, root)
	//t.Log("\\")
	//result := removeLeafNodes(root, 1)
	//InOrderTraversal(t, result)

	//root := &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 2}
	//root.Right = &TreeNode{Val: 3}
	//
	//InOrderTraversal(t, root)
	//t.Log("\\")
	//result := removeLeafNodes(root, 1)
	//InOrderTraversal(t, result)

	//root := &TreeNode{Val: 1}
	//root.Left = &TreeNode{Val: 3}
	//root.Right = &TreeNode{Val: 3}
	//root.Left.Left = &TreeNode{Val: 3}
	//root.Left.Right = &TreeNode{Val: 2}
	//
	//InOrderTraversal(t, root)
	//t.Log("\\")
	//result := removeLeafNodes(root, 3)
	//InOrderTraversal(t, result)

	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 2}
	root.Left.Left.Left = &TreeNode{Val: 2}

	InOrderTraversal(t, root)
	t.Log("\\")
	result := removeLeafNodes(root, 2)
	InOrderTraversal(t, result)

}

func InOrderTraversal(t *testing.T, root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversal(t, root.Left)
	t.Log(root.Val)
	InOrderTraversal(t, root.Right)
}
