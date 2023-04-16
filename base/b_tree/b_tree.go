package b_tree

type Node struct {
	Keys   []int
	Child  []Node
	Leaf   bool
	T      int
	keyCnt int
}

func NewNode(t int) Node {
	return Node{
		Keys:   make([]int, 2*t-1),
		Child:  make([]Node, 2*t),
		Leaf:   false,
		T:      t,
		keyCnt: 0,
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
	for i < x.keyCnt && key < x.Keys[i] {
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
	i := x.keyCnt - 1

	if x.Leaf {
		for ; i >= 0 && k < x.Keys[i]; i-- {
			x.Keys[i+1] = x.Keys[i]
		}

		x.Keys[i+1] = k
		x.keyCnt++

		return
	}

	for i >= 0 && k < x.Keys[i] {
		i--
	}

	i++
	if x.Child[i].keyCnt == 2*x.T-1 {
		x.split(i)
		if k > x.Keys[i] {
			i++
		}
	}

	x.Child[i].insert(k)
}

func (x *Node) split(pos int) {
	t := x.T
	y, z := &x.Child[pos], NewNode(t)
	z.Leaf = y.Leaf
	z.keyCnt = t - 1
	for j := 0; j < t-1; j++ {
		z.Keys[j] = y.Keys[j+t]
	}

	if !y.Leaf {
		for j := 0; j < t; j++ {
			z.Child[j] = y.Child[j+t]
		}
	}
	y.keyCnt = t - 1

	for j := x.keyCnt; j >= pos+1; j-- {
		x.Child[j+1] = x.Child[j]
	}
	x.Child[pos+1] = z

	for j := x.keyCnt - 1; j >= pos; j-- {
		x.Keys[j+1] = x.Keys[j]
	}

	x.Keys[pos] = y.Keys[t-1]
	x.keyCnt++
}
