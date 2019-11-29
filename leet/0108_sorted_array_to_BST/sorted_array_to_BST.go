package _0108_sorted_array_to_BST

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

func PreOrder(x *TreeNode, list *[]int) {
	if x != nil {
		*list = append(*list, x.Val)
		PreOrder(x.Left, list)
		PreOrder(x.Right, list)
	}
}

func LevelOrder(x *TreeNode) []int {
	list := make([]int, 0)
	level := 1
	for traversalLevel(x, level, &list) {
		level++
	}
	return list
}

func traversalLevel(x *TreeNode, level int, list *[]int) bool {
	if x == nil {
		return false
	}

	if level == 1 {
		*list = append(*list, x.Val)
		return true
	}

	left := traversalLevel(x.Left, level-1, list)
	right := traversalLevel(x.Right, level-1, list)

	return left || right
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
