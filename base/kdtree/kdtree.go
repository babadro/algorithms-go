package kdtree

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

func (t *KDTree) Insert(point []int) {
	t.Root = t.insertHelper(t.Root, point, 0)
}

func pointsAreEqual(p1, p2 []int) bool {
	for i := range p1 {
		if p1[i] != p2[i] {
			return false
		}
	}

	return true
}

func (t *KDTree) searchHelper(node *Node, point []int, depth uint) *Node {
	if node == nil {
		return nil
	}

	if pointsAreEqual(node.Point, point) {
		return node
	}

	cd := int(depth) % t.K

	if point[cd] < node.Point[cd] {
		return t.searchHelper(node.Left, point, depth+1)
	}

	return t.searchHelper(node.Right, point, depth+1)
}

func (t *KDTree) Search(point []int) *Node {
	return t.searchHelper(t.Root, point, 0)
}

func (t *KDTree) findMinHelper(node *Node, d int, depth uint) *Node {
	if node == nil {
		return nil
	}

	cd := int(depth) % t.K

	if cd == d {
		if node.Left == nil {
			return node
		}

		return t.findMinHelper(node.Left, d, depth+1)
	}

	return minNode(node,
		t.findMinHelper(node.Left, d, depth+1),
		t.findMinHelper(node.Right, d, depth+1), d)
}

// todo 2 unit tests
func (t *KDTree) FindMin(d int) *Node {
	return t.findMinHelper(t.Root, d, 0)
}

func minNode(x, y, z *Node, d int) *Node {
	res := x
	if y != nil && y.Point[d] < res.Point[d] {
		res = y
	}

	if z != nil && z.Point[d] < res.Point[d] {
		res = z
	}

	return res
}

// todo 2 doesn't work
// may be it'll help
// https://github.com/kyroy/kdtree/blob/master/kdtree.go
func (t *KDTree) DeleteNode(point []int) *Node {
	return t.deleteNodeHelper(t.Root, point, 0)
}

func (t *KDTree) deleteNodeHelper(node *Node, point []int, depth uint) *Node {
	if node == nil {
		return nil
	}

	cd := int(depth) % t.K

	if pointsAreEqual(node.Point, point) {
		if node.Right != nil {
			min := t.findMinHelper(node.Right, cd, 0) // почему не depth+1 ?

			copyPoint(node.Point, min.Point)

			node.Right = t.deleteNodeHelper(node.Right, min.Point, depth+1)
		} else if node.Left != nil {
			min := t.findMinHelper(node.Left, cd, 0) //почему не depth+1 ?

			copyPoint(node.Point, min.Point)

			node.Right = t.deleteNodeHelper(node.Left, min.Point, depth+1) // почсему node.Right, а не node.Left?
		} else { // Leaf
			return nil
		}

		return node
	}

	if point[cd] < node.Point[cd] {
		node.Left = t.deleteNodeHelper(node.Left, point, depth+1)
	} else {
		node.Right = t.deleteNodeHelper(node.Right, point, depth+1)
	}

	return node
}

func copyPoint(dst, src []int) {
	for i := range src {
		dst[i] = src[i]
	}
}

func InOrderFunc(node *Node, f func(*Node)) {
	if node != nil {
		InOrderFunc(node.Left, f)
		f(node)
		InOrderFunc(node.Right, f)
	}
}
