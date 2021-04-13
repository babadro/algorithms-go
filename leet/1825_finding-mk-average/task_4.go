package _1825_finding_mk_average

import "sort"

// todo 1 https://leetcode.com/problems/finding-mk-average/discuss/1155397/Java%3A-Binary-Search-(copied-from-Arrays-Class)-%2B-ArrayList-as-SlidingWindow
type MKAverage struct {
	m       int
	k       int
	resLen  int
	nums    []int
	indexes []int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{
		m:      m,
		k:      k,
		resLen: m - 2*k,
	}
}

func (this *MKAverage) slideByRemovingElement() {
	n, m := len(this.nums), this.m
	if n >= m {
		i := n - m
		copy(this.indexes[i:len(this.indexes)-1], this.indexes[i+1:])
		this.indexes = this.indexes[:len(this.indexes)-1]
	}
}

func (this *MKAverage) AddElement(num int) {
	this.arr = append(this.arr, num)
}

func (this *MKAverage) CalculateMKAverage() int {
	if len(this.arr) < this.m {
		return -1
	}

	cur := make([]int, this.m)
	copy(cur, this.arr[len(this.arr)-this.m:])
	sort.Ints(cur)
	cur = cur[this.k : len(cur)-this.k]

	sum := 0
	for _, num := range cur {
		sum += num
	}

	return sum / len(cur)
}
