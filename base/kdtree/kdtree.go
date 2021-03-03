package kdtree

import "math"

// todo base delete func (file temp3) and unit tests
type Node struct {
	Point       []int
	Left, Right *Node
}

type KDTree struct {
	K    int
	Root *Node
}

func NewNode(point []int) *Node {
	return &Node{Point: point}
}
func (t *KDTree) insertHelper(node *Node, point []int, depth uint) *Node {
	if node == nil {
		return NewNode(point)
	}

	cd := int(depth) % t.K

	if point[cd] < node.Point[cd] {
		node.Left = t.insertHelper(node.Left, point, depth+1)
	} else {
		node.Right = t.insertHelper(node.Right, point, depth+1)
	}

	return node
}

func (t *KDTree) Insert(root *Node, point []int) *Node {
	return t.insertHelper(root, point, 0)
}

func pointsAreEqual(p1, p2 []int) bool {
	for i := range p1 {
		if p1[i] != p2[i] {
			return false
		}
	}

	return true
}

func (t *KDTree) searchHelper(node *Node, point []int, depth uint) bool {
	if node == nil {
		return false
	}

	if pointsAreEqual(node.Point, point) {
		return true
	}

	cd := int(depth) % t.K

	if point[cd] < node.Point[cd] {
		return t.searchHelper(node.Left, point, depth+1)
	}

	return t.searchHelper(node.Right, point, depth+1)
}

func (t *KDTree) Search(point []int) bool {
	return t.searchHelper(t.Root, point, 0)
}

func (t *KDTree) findMinHelper(node *Node, d int, depth uint) int {
	if node == nil {
		return math.MaxInt64
	}

	cd := int(depth) % t.K

	if cd == d {
		if node.Left == nil {
			return node.Point[d]
		}

		return min2(node.Point[d], t.findMinHelper(node.Left, d, depth+1))
	}

	return min3(node.Point[d],
		t.findMinHelper(node.Left, d, depth+1),
		t.findMinHelper(node.Right, d, depth+1))
}

func (t *KDTree) FindMin(d int) int {
	return t.findMinHelper(t.Root, d, 0)
}

func min3(a, b, c int) (res int) {
	return min2(a, min2(b, c))
}

func min2(a, b int) int {
	if a < b {
		return a
	}

	return b
}
