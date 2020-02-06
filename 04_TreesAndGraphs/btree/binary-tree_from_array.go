package btree

import "math"

const (
	null = math.MaxInt64
)

type level struct {
	nodes  []*TreeNode
	length int
}

//[1,1]
//[1,2,3,4,5,6]
//[2,3,9,10,7,8,6,5,4,11,1]
//[1,null,2,3,4,null,null,5,6]

func ArrayToBinaryTree(arr []int) (root *TreeNode) {
	if len(arr) == 0 {
		return nil
	}
	root = &TreeNode{Val: arr[0]}
	levels := make([]level, 2)
	levels[0] = level{nodes: []*TreeNode{root}, length: 1}
	i, currLevel := 0, 1
	for i < len(arr) {
		levelLen := levels[currLevel-1].length * 2
		for i < levelLen && i < len(arr) {
			var newNode *TreeNode
			if arr[i] != null {
				newNode = &TreeNode{Val: arr[i]}
				levels[currLevel].length++
			} else {
				newNode = nil
			}
			levels[currLevel].nodes = append(levels[currLevel].nodes, newNode)
			i++
		}
		levels = append(levels, level{})
		currLevel++
	}

	for k, level := range levels[:len(levels)-1] {
		j := 0
		for _, node := range level.nodes {
			if node != nil {
				node.Left = levels[k+1].nodes[j]
				j++
				node.Right = levels[k+1].nodes[j]
				j++
			}
		}
	}

	return root
}
