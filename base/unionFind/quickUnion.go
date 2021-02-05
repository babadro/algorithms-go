// QuickUnion is also too slow
package unionFind

type QuickUnion struct {
	id []int
}

func NewQuickUnion(n int) QuickUnion {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return QuickUnion{id: id}
}

func (qu *QuickUnion) root(i int) int {
	for i != qu.id[i] {
		i = qu.id[i]
	}

	return i
}

func (qu *QuickUnion) Find(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu *QuickUnion) Union(p, q int) {
	rootP := qu.root(p)
	rootQ := qu.root(q)
	qu.id[rootP] = rootQ
}
