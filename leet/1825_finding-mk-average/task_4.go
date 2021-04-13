package _1825_finding_mk_average

import "sort"

// todo 1. TLE
type MKAverage struct {
	arr  []int
	m, k int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{m: m, k: k}
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

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
