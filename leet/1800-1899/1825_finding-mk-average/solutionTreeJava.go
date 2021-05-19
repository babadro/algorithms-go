package _1825_finding_mk_average

import "container/list"

// https://leetcode.com/problems/finding-mk-average/discuss/1155190/Segment-Treejava-68ms
// passed. good solution. todo 2 need to understand
type MKAverageJava struct {
	queue *list.List
	count []int
	sum   []int
	m, k  int
}

func ConstructorJava(m int, k int) MKAverageJava {
	return MKAverageJava{
		queue: list.New(),
		count: make([]int, 400_001),
		sum:   make([]int, 400_001),
		m:     m,
		k:     k,
	}
}

func (mk *MKAverageJava) AddElement(num int) {
	if mk.queue.Len() == mk.m {
		v := mk.queue.Front().Value.(int)
		mk.queue.Remove(mk.queue.Front())

		mk.insert(1, 0, 100_000, v, -1)
	}

	mk.insert(1, 0, 100_000, num, 1)
	mk.queue.PushBack(num)
}

func (mk *MKAverageJava) CalculateMKAverage() int {
	if mk.queue.Len() < mk.m {
		return -1
	}

	s := mk.k + 1
	e := mk.m - mk.k

	return mk.query(1, 0, 100_000, s, e) / (mk.m - 2*mk.k)
}

func (mk *MKAverageJava) insert(node, l, r, v, d int) {
	mk.count[node] += d
	mk.sum[node] += d * v
	if l == r {
		return
	}

	m := (l + r) / 2
	if v <= m {
		mk.insert(node*2, l, m, v, d)
	} else {
		mk.insert((node*2)+1, m+1, r, v, d)
	}
}

func (mk *MKAverageJava) query(node, l, r, s, e int) int {
	if l == r {
		c := e - s + 1

		return c * l
	} else if mk.count[node] == e-s+1 {
		return mk.sum[node]
	} else {
		m := (l + r) / 2
		c1 := mk.count[node*2]

		if c1 >= e {
			return mk.query(node*2, l, m, s, e)
		} else if c1 >= s {
			return mk.query(node*2, l, m, s, c1) + mk.query((node*2)+1, m+1, r, 1, e-c1)
		} else { // c1 < s
			return mk.query((node*2)+1, m+1, r, s-c1, e-c1)
		}
	}
}
