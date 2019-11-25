package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func (t *Tree) Insert(z *TreeNode) {
	var parent *TreeNode = nil
	child := t.Root
	for child != nil {
		parent = child
		if z.Val < child.Val {
			child = child.Left
		} else {
			child = child.Right
		}
	}
	if parent == nil {
		t.Root = z
	} else if z.Val < parent.Val {
		parent.Left = z
	} else {
		parent.Right = z
	}
}

func main() {
	sortedArrayToBST([]int{1, 3})
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	middle := len(nums) / 2
	left := nums[:middle]
	right := nums[middle:]
	root := TreeNode{Val: right[0]}
	right = right[1:]
	root.Left = sortedArrayToBST(left)
	root.Right = sortedArrayToBST(right)
	return &root
}
