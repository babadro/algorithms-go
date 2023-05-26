package _1101_the_earliest_moment_when_everyone_become_friends

import (
	"sort"
)

// #bnsrg
func earliestAcq(logs [][]int, n int) int {
	sort.Slice(logs, func(i, j int) bool {
		return logs[i][0] < logs[j][0]
	})

	uf := newUnionFind(n)
	for _, log := range logs {
		if uf.find(log[1], log[2]) {
			continue
		}

		if uf.union(log[1], log[2]) == n {
			return log[0]
		}
	}

	return -1
}

type unionFind struct {
	id, size []int
}

func newUnionFind(n int) unionFind {
	id, size := make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		id[i], size[i] = i, 1
	}

	return unionFind{id, size}
}

func (uf *unionFind) root(i int) int {
	for i != uf.id[i] {
		uf.id[i] = uf.id[uf.id[i]]
		i = uf.id[i]
	}

	return i
}

func (uf *unionFind) find(i, j int) bool {
	return uf.root(i) == uf.root(j)
}

func (uf *unionFind) union(p, q int) (size int) {
	i, j := uf.root(p), uf.root(q)

	if uf.size[i] < uf.size[j] {
		uf.size[j] += uf.size[i]
		uf.id[i] = uf.id[j]

		return uf.size[j]
	} else {
		uf.size[i] += uf.size[j]
		uf.id[j] = uf.id[i]

		return uf.size[i]
	}
}
