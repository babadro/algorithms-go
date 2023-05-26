// https://www.cs.princeton.edu/~rs/AlgsDS07/01UnionFind.pdf
// dyx, tptl?
package unionFind

// todo base unit tests
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

func (qu *WQUPC) Root(i int) int {
	for i != qu.id[i] {
		qu.id[i] = qu.id[qu.id[i]] // path compression todo base understand
		i = qu.id[i]
	}

	return i
}

func (qu *WQUPC) Find(p, q int) bool {
	return qu.Root(p) == qu.Root(q)
}

func (qu *WQUPC) Union(p, q int) {
	i := qu.Root(p)
	j := qu.Root(q)

	if i == j {
		return
	}

	if qu.size[i] < qu.size[j] {
		qu.id[i] = qu.id[j]
		qu.size[j] += qu.size[i]
	} else {
		qu.id[j] = qu.id[i]
		qu.size[i] += qu.size[j]
	}
}
