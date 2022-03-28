package _2213_longest_subsring_of_one_repeating_character

// passed https://leetcode.com/problems/longest-substring-of-one-repeating-character/discuss/1879490/Segment-Tree
func longestRepeating2(s string, queryCharacters string, queryIndices []int) []int {
	var res []int
	powOf2, sz := 1, len(s)

	for powOf2 < sz {
		powOf2 <<= 1
	}

	st := make([]node2, powOf2*2)
	for i := range st {
		st[i].sz = 1 // initial value
	}

	for i := 0; i < len(s); i++ {
		stSet(st, i, s[i], 0, 0, powOf2-1)
	}

	for j := 0; j < len(queryCharacters); j++ {
		res = append(res, stSet(st, queryIndices[j], queryCharacters[j], 0, 0, powOf2-1))
	}

	return res
}

type node2 struct {
	lc, rc                 byte
	pref, suf, longest, sz int
}

func (n *node2) merge(l, r node2) {
	n.longest = max(l.longest, r.longest)
	if l.rc == r.lc {
		n.longest = max(n.longest, l.suf+r.pref)
	}

	n.sz = l.sz + r.sz
	n.lc, n.rc = l.lc, r.rc

	if l.pref == l.sz && l.lc == r.lc {
		n.pref = l.pref + r.pref
	} else {
		n.pref = l.pref
	}

	if r.suf == r.sz && r.rc == l.rc {
		n.suf = r.suf + l.suf
	} else {
		n.suf = r.suf
	}
}

func stSet(st []node2, pos int, ch byte, i, l, r int) int {
	if l <= pos && pos <= r {
		if l != r {
			m := l + (r-l)/2
			li := 2*i + 1
			ri := 2*i + 2

			stSet(st, pos, ch, li, l, m)
			stSet(st, pos, ch, ri, m+1, r)
			st[i].merge(st[li], st[ri])
		} else {
			st[i].lc, st[i].rc = ch, ch
			st[i].suf, st[i].pref, st[i].longest = 1, 1, 1
		}
	}

	return st[i].longest
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

// todo 2 implement it:
// https://leetcode.com/problems/longest-substring-of-one-repeating-character/discuss/1877448/Merge-intervals-using-map
// notes: use bst and map for replace c++ map. Use map and maxHeap for replace c++ multiset. Implement bst with generics.
