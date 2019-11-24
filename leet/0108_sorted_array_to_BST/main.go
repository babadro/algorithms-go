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
	var tree Tree
	middle := len(nums) / 2
	left := nums[:middle]
	right := nums[middle:]
	tree.Insert(&TreeNode{Val: right[0]})
	right = right[1:]
	i := middle - 1
	for _, rNum := range right {
		tree.Insert(&TreeNode{Val: rNum})
		if i >= 0 {
			tree.Insert(&TreeNode{Val: left[i]})
			i--
		}
	}
	if i == 0 {
		tree.Insert(&TreeNode{Val: left[0]})
	}
	return tree.Root
}
