package _128_longest_consecutive_sequence

type uf struct {
	ids  []int
	size []int
}

func newUF(n int) uf {
	res := uf{
		ids:  make([]int, n),
		size: make([]int, n),
	}

	for i := range res.ids {
		res.ids[i] = i
		res.size[i] = 1
	}

	return res
}

func (u *uf) root(i int) int {
	for i != u.ids[i] {
		u.ids[i] = u.ids[u.ids[i]]
		i = u.ids[i]
	}

	return i
}

func (u *uf) union(i, j int) {
	i, j = u.root(i), u.root(j)
	if i == j {
		return
	}

	if u.size[i] > u.size[j] {
		u.ids[j] = i
		u.size[i] += u.size[j]
	} else {
		u.ids[i] = j
		u.size[j] += u.size[i]
	}
}

func longestConsecutive(nums []int) int {
	n := 0
	valToIDx := make(map[int]int)
	for _, num := range nums {
		if _, ok := valToIDx[num]; !ok {
			valToIDx[num] = n
			n++
		}
	}

	u := newUF(n)

	for val, idx := range valToIDx {
		if upIDx, ok := valToIDx[val+1]; ok {
			u.union(idx, upIDx)
		} else if downIDx, ok := valToIDx[val-1]; ok {
			u.union(idx, downIDx)
		}
	}

	res := 0
	for _, size := range u.size {
		res = max(res, size)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
