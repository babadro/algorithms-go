package _2101_detonate_the_maximum_bombs

// union find approach doesn't work here. Explanation why:
// https://leetcode.com/problems/detonate-the-maximum-bombs/solutions/1625917/union-find-146-160-test-cases-passed-why/
func maximumDetonationUF(bombs [][]int) int {
	uf := newUnionFind(len(bombs))
	for i := range bombs {
		for j := i + 1; j < len(bombs); j++ {
			if uf.root(i) == uf.root(j) {
				continue
			}

			small, big := bombs[i], bombs[j]
			if small[2] > big[2] {
				small, big = big, small
			}

			if distance(small, big) <= big[2]*big[2] {
				uf.add(i, j)
			}
		}
	}

	chainsOfBomb, res := make(map[int]int), 1
	for i := range uf.id {
		if root := uf.root(i); root != i {
			chainsOfBomb[root]++
			res = max(res, chainsOfBomb[root]+1)
		}
	}

	return res
}

type unionFind struct {
	id, weight []int
}

func newUnionFind(capacity int) unionFind {
	id, weight := make([]int, capacity), make([]int, capacity)
	for i := range id {
		id[i], weight[i] = i, 1
	}

	return unionFind{id, weight}
}

func (uf *unionFind) root(i int) int {
	for i != uf.id[i] {
		uf.id[i] = uf.id[uf.id[i]]
		i = uf.id[i]
	}

	return i
}

func (uf *unionFind) add(i, j int) {
	i, j = uf.root(i), uf.root(j)

	if uf.weight[i] < uf.weight[j] {
		uf.weight[j] += uf.weight[i]
		uf.id[i] = uf.id[j]
	} else {
		uf.weight[i] += uf.weight[j]
		uf.id[j] = uf.id[i]
	}
}
