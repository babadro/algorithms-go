package _1325_delete_leaves_with_a_given_value

//https://leetcode.com/contest/weekly-contest-172/problems/delete-leaves-with-a-given-value/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	stack := make([]ExtendedTreeNode, 0)
	queue := make([]ExtendedTreeNode, 0)

	queue = append(queue, ExtendedTreeNode{root, nil})
	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		stack = append(stack, n)

		if n.TreeNode.Right != nil {
			queue = append(queue, ExtendedTreeNode{n.TreeNode.Right, n.TreeNode})
		}
		if n.TreeNode.Left != nil {
			queue = append(queue, ExtendedTreeNode{n.TreeNode.Left, n.TreeNode})
		}
	}

	for i := len(stack) - 1; i >= 0; i-- {
		if stack[i].TreeNode.Val == target && stack[i].TreeNode.Left == nil && stack[i].TreeNode.Right == nil {
			parent := stack[i].Parent
			if parent == nil {
				return nil
			}
			if parent.Right == stack[i].TreeNode {
				parent.Right = nil
			} else {
				parent.Left = nil
			}
		}
	}

	return root
}

type ExtendedTreeNode struct {
	TreeNode *TreeNode
	Parent   *TreeNode
}

func removeLeafNodesRecursive(root *TreeNode, target int) *TreeNode {
	if root.Left != nil {
		root.Left = removeLeafNodes(root.Left, target)
	}
	if root.Right != nil {
		root.Right = removeLeafNodes(root.Right, target)
	}
	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}
