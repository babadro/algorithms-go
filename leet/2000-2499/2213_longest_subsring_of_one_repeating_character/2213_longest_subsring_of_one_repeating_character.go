package _2213_longest_subsring_of_one_repeating_character

// https://leetcode.com/problems/longest-substring-of-one-repeating-character/discuss/1866110/Java-clean-Segment-tree-solution
// todo 1
func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {

}

type node struct {
	max             int
	prefSt, prefEnd int
	suffSt, suffEnd int
}

func newNode(m, prefSt, prefEnd, suffSt, suffEnd int) node {
	return node{
		max:     m,
		prefSt:  prefSt,
		prefEnd: prefEnd,
		suffSt:  suffSt,
		suffEnd: suffEnd,
	}
}

type segmentTree struct {
	tree []node
	s    []byte
}

func (st *segmentTree) merge(left, right node, tl, tm, tr int) node {
	maximum := max(left.max, right.max)
	prefSt, prefEnd := left.prefSt, left.prefEnd
	suffSt, suffEnd := right.suffSt, right.suffEnd

	if st.s[tm] == st.s[tm+1] {
		maximum = max(maximum, right.prefEnd-left.suffSt+1)
		if left.prefEnd-left.prefSt+1 == tm-tl+1 {
			prefEnd = right.prefEnd
		}
		if right.suffEnd-right.suffEnd+1 == tr-tm {
			suffSt = left.suffSt
		}
	}

	return newNode(maximum, prefSt, prefEnd, suffSt, suffEnd)
}

func (st *segmentTree) build(pos, tl, tr int) {
	if tl == tr {
		st.tree[pos] = newNode(1, tl, tl, tr, tr)
	} else {
		tm := tl + (tr-tl)/2
		st.build(2*pos+1, tl, tm)
		st.build(2*pos+2, tm+1, tr)

		st.tree[pos] = st.merge(st.tree[2*pos+1], st.tree[2*pos+2], tl, tm, tr)
	}
}
