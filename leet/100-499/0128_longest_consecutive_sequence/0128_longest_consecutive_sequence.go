package _128_longest_consecutive_sequence

// tptl. passed
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
	iRoot, jRoot := u.root(i), u.root(j)
	if iRoot == jRoot {
		return
	}

	if u.size[iRoot] > u.size[jRoot] {
		u.ids[iRoot] = jRoot
		u.size[jRoot] += u.size[iRoot]
	} else {
		u.ids[jRoot] = iRoot
		u.size[iRoot] += u.size[jRoot]
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

// tptl. passed. best solution. elegant and fast
func longestConsecutive2(nums []int) int {
	numToSeqLen, ans := make(map[int]int), 0

	for _, num := range nums {
		if _, ok := numToSeqLen[num]; ok {
			continue
		}

		nextLen, nextOk := numToSeqLen[num+1]
		prevLen, prevOk := numToSeqLen[num-1]

		length := nextLen + prevLen + 1
		numToSeqLen[num] = length

		if nextOk {
			numToSeqLen[num+nextLen] = length
		}

		if prevOk {
			numToSeqLen[num-prevLen] = length
		}

		ans = max(ans, length)
	}

	return ans
}
