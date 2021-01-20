package b_tree

// todo 1: insertion, deletion, unit tests
type BTreeNode struct {
	keys []int
	t    int
	C    []BTreeNode
	n    int
	leaf bool
}

func New(t int, leaf bool) BTreeNode {
	return BTreeNode{
		keys: make([]int, 2*t-1),
		t:    t,
		C:    make([]BTreeNode, 2*t),
		leaf: false,
		n:    0,
	}
}

func (b *BTreeNode) Traverse(f func(n *BTreeNode)) {
	i := 0
	for ; i < b.n; i++ {
		if !b.leaf {
			b.C[i].Traverse(f)
		}

		f(b)
	}

	if !b.leaf {
		b.C[i].Traverse(f)
	}
}

func (b *BTreeNode) Search(k int) *BTreeNode {
	i := 0
	for i < b.n && k > b.keys[i] {
		i++
	}

	if b.keys[i] == k {
		return b
	}

	if b.leaf {
		return nil
	}

	return b.C[i].Search(k)
}
