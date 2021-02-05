// QuickFind is slow!
package unionFind

type QuickFind struct {
	id []int
}

func NewQuckFind(n int) QuickFind {
	id := make([]int, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}

	return QuickFind{id: id}
}

func (qf *QuickFind) Find(p, q int) bool {
	return qf.id[p] == qf.id[q]
}

func (qf *QuickFind) Union(p, q int) {
	pid := qf.id[p]

	for i := 0; i < len(qf.id); i++ {
		if qf.id[i] == pid {
			qf.id[i] = qf.id[q]
		}
	}
}
