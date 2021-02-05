// https://www.cs.princeton.edu/~rs/AlgsDS07/01UnionFind.pdf
// dyx, tptl?
package unionFind

// todo 1 unit tests
// WQUPC - weighted quick-union with path compression
type WQUPC struct {
	id   []int
	size []int
}

func NewWQUPC(n int) WQUPC {
	id, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
		size[i] = 1
	}

	return WQUPC{
		id:   id,
		size: size,
	}
}

func (qu *WQUPC) root(i int) int {
	for i != qu.id[i] {
		qu.id[i] = qu.id[qu.id[i]] // path compression todo 1 understand
		i = qu.id[i]
	}

	return i
}

func (qu *WQUPC) Find(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu *WQUPC) Union(p, q int) {
	i := qu.root(p)
	j := qu.root(q)

	if qu.size[i] < qu.size[j] {
		qu.id[i] = qu.id[j]
		qu.size[j] += qu.size[i]
	} else {
		qu.id[j] = qu.id[i]
		qu.size[i] += qu.size[j]
	}
}
