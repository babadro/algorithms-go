package binaryTree

import "math"

const (
	Null = math.MaxInt64
)

type level struct {
	nodes  []*Node
	length int
}

//[1,2]
//[1,2,3,4,5,6]
//[2,3,9,10,7,8,6,5,4,11,1]
//[1,null,2,3,4,null,null,5,6]
// TODO 2 infinitive loop with []int{1, 2, 3, 4, 5, 6, 7, binaryTree.Null, binaryTree.Null, 8, 9, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, binaryTree.Null, 10}
func ArrayToBinaryTree(arr []int) (root *Node) {
	if len(arr) == 0 {
		return nil
	}
	root = &Node{Val: arr[0]}
	levels := make([]level, 1)
	levels[0] = level{nodes: []*Node{root}, length: 1}
	i, currLevel := 1, 1
	for {
		levels = append(levels, level{})
		levelLen := levels[currLevel-1].length * 2
		endLevel := i + levelLen
		for i < endLevel && i < len(arr) {
			var newNode *Node
			if arr[i] != Null {
				newNode = &Node{Val: arr[i]}
				levels[currLevel].length++
			}
			levels[currLevel].nodes = append(levels[currLevel].nodes, newNode)
			i++
		}
		currLevel++
		if i == len(arr) {
			break
		}
	}

	for k, level := range levels[:len(levels)-1] {
		j := 0
		nextLevelNodes := levels[k+1].nodes
		for _, node := range level.nodes {
			if node != nil {
				if j == len(nextLevelNodes) {
					break
				}
				node.Left = nextLevelNodes[j]
				j++
				if j == len(nextLevelNodes) {
					break
				}
				node.Right = nextLevelNodes[j]
				j++
			}
		}
	}

	return root
}
