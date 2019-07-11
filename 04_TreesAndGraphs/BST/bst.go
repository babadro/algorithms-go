package bst

type Node struct {
	Key                 int
	Left, Right, Parent *Node
}

type Tree struct {
	Root *Node
}

func Inorder(x *Node, list *[]int) {
	if x != nil {
		Inorder(x.Left, list)
		*list = append(*list, x.Key)
		Inorder(x.Right, list)
	}
}

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

func Search(x *Node, k int) *Node {
	if x == nil || k == x.Key {
		return x
	}
	if k < x.Key {
		return Search(x.Left, k)
	}
	return Search(x.Right, k)
}

func IterativeSearch(x *Node, k int) *Node {
	for x != nil && k != x.Key {

	}

}

/* pseudocode

InorderTreeWalk(x)
	if x != NIL
		InorderTreeWalk(x.left)
		print x.key
		InorderTreeWalk(x.right)

TreeSearch(x, k) // x - root дерева
	if x == NIL or k == x.key
		return x
	if k < x.key
		return TreeSearch(x.left, k)
	return TreeSearch(x.right, k)


IterativeTreeSearch(x, k)
	while x != NIL and k != x.key
		if k < x.key
			x = x.left
		else
			x = x.right
	return x

TreeMinimum(x)
	while x.left != NIL
		x = x.left
	return x


TreeMaximum(x)
	while x.right != NIL
		x = x.right
	return x

TreeSuccessor(x)
	if x.right != NIL
		return TreeMinimum(x.right)
	y = x.p
	while y != NIL and x == y.right // тут именно сравнение указателей, а не ключей. Т.е. х должен быть правым листом y, чтобы цикл продолжался.
		x = y
		y = y.p
	return y

TreeInsert(T, z)
	y = NIL
	x = T.root
	while x != NIL
		y = x
		if z.key < x.key
			x = x.left
		else
			x = x.right
	z.p = y
	if y == NIL
		T.root = z
	else if z.key < y.key
		y.left = z
	else
		y.right = z

Transplant(T, u, v)
	if u.p == NIL
		T.root = v
	elseif u == u.p.left
		u.p.left = v
	else
		u.p.right = v
	if v != NIL
		v.p = u.p


TreeDelete(T, z)
	if z.left == NIL
		Transplant(T, z, z.right)
	else if z.right == NIL
		Transplant(T, z, z.left)
	else
		y = TreeMinimum(z.right)
		if y.p != z
			Transplant(T, y, y.right)
			y.right = z.right
			y.right.p = y
		Transplant(T, z, y)
		y.left = z.left
		y.left.p = y
*/
