package b_tree

type Node struct {
	Keys  []int
	Child []Node
	Leaf  bool
	T     int
	n     int
}

func NewNode(t int) Node {
	return Node{
		Keys:  make([]int, 0),
		Child: make([]Node, 0),
		Leaf:  false,
		T:     t,
		n:     0,
	}
}

func NewBTree(t int) Node {
	root := NewNode(t)
	root.Leaf = true

	return root
}

func (x *Node) Search(key int) *Node {
	if x == nil {
		return nil
	}

	i := 0
	for i < x.n && key < x.Keys[i] {
		i++
	}

	if key == x.Keys[i] {
		return x
	}

	if x.Leaf {
		return nil
	}

	return x.Child[i].Search(key)
}

func (x *Node) insert(k int) {
	i := x.n - 1

	if x.Leaf {
		for ; i >= 0 && k < x.Keys[i]; i-- {
			x.Keys[i+1] = x.Keys[i]
		}

		x.Keys[i+1] = k
		x.n++

		return
	}

	for i >= 0 && k < x.Keys[i] {
		i--
	}

	i++
	tmp := x.Child[i]
	if tmp.n == 2*tmp.T-1 {
		// split todo
	}

	x.Child[i].insert(k)
}

func (x *Node) split(pos int) {
	t := x.T
	y, z := &x.Child[pos], NewNode(t)
	z.Leaf = y.Leaf
	z.n = t - 1
	for j := 0; j < t-1; j++ {
		z.Keys[j] = y.Keys[j+t]
	}

	if !y.Leaf {
		for j := 0; j < t; j++ {
			z.Child[j] = y.Child[j+t]
		}
	}

	y.n = t - 1
	for j := x.n; j >= pos+1; j-- {
		x.Child[j+1] = x.Child[j]
	}
	x.Child[pos+1] = z

	for j := x.n - 1; j >= pos; j-- {
		x.Keys[j+1] = x.Keys[j]
	}

	x.Keys[pos] = y.Keys[t-1]
	x.n = x.n + 1
}
