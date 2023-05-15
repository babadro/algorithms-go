package binaryTree

import "math"

const (
	Null = math.MaxInt64
)

// [1,2]
// [1,2,3,4,5,6]
// [2,3,9,10,7,8,6,5,4,11,1]
// [1,null,2,3,4,null,null,5,6]
func ArrayToBinaryTree(nums []int) *Node {
	if nums[0] == Null {
		return nil
	}

	root := &Node{Val: nums[0]}
	queue := []*Node{root}
	i := 1
	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]
		if nums[i] != Null {
			node.Left = &Node{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nums) && nums[i] != Null {
			node.Right = &Node{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func InterfacesArrToBinaryTree(nums []interface{}) *Node {
	if nums[0] == nil {
		return nil
	}

	root := &Node{Val: nums[0].(int)}
	queue := []*Node{root}
	i := 1
	for i < len(nums) {
		node := queue[0]
		queue = queue[1:]
		if nums[i] != nil {
			node.Left = &Node{Val: nums[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nums) && nums[i] != nil {
			node.Right = &Node{Val: nums[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
