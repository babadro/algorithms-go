package bst

// todo base implement with generics
type Node struct {
	Key                 int
	Left, Right, Parent *Node
}

type Tree struct {
	Root *Node
}

func Inorder(x *Node, f func(node *Node)) {
	if x != nil {
		Inorder(x.Left, f)
		f(x)
		Inorder(x.Right, f)
	}
}

// dyx
func (t *Tree) Insert(z *Node) {
	var parent *Node = nil
	child := t.Root
	for child != nil {
		parent = child
		if z.Key < child.Key {
			child = child.Left
		} else {
			child = child.Right
		}
	}
	z.Parent = parent
	if parent == nil {
		t.Root = z
	} else if z.Key < parent.Key {
		parent.Left = z
	} else {
		parent.Right = z
	}
}

func (t *Tree) transplant(u, v *Node) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}
	if v != nil {
		v.Parent = u.Parent
	}
}

// todo base
func (t *Tree) Delete(z *Node) {
	if z.Left == nil {
		t.transplant(z, z.Right)
	} else if z.Right == nil {
		t.transplant(z, z.Left)
	} else {
		y := Min(z.Right)
		if y != z.Right {
			t.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.Parent = y
		}
		t.transplant(z, y)
		y.Left = z.Left
		y.Left.Parent = y
	}
}

/* // Recursive Search
func Search(x *Node, k int) *Node {
	if x == nil || k == x.Key {
		return x
	}
	if k < x.Key {
		return Search(x.Left, k)
	}
	return Search(x.Right, k)
}
*/

// dyx
func Search(x *Node, k int) *Node {
	for x != nil && k != x.Key {
		if k < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	return x
}

// dyx
func Min(x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

// dyx
func Max(x *Node) *Node {
	for x.Right != nil {
		x = x.Right
	}
	return x
}

func Successor(x *Node) *Node {
	if x.Right != nil {
		return Min(x.Right)
	}
	y := x.Parent
	for y != nil && x == y.Right {
		x = y
		y = y.Parent
	}
	return y
}
