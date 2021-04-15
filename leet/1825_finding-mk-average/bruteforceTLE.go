package _1825_finding_mk_average

import "sort"

// correct, but got TLE. bruteforce solution
type MKAverage1 struct {
	arr  []int
	m, k int
}

func Constructor1(m int, k int) MKAverage1 {
	return MKAverage1{m: m, k: k}
}

func (this *MKAverage1) AddElement(num int) {
	this.arr = append(this.arr, num)
}

func (this *MKAverage1) CalculateMKAverage() int {
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
 * Your MKAverage1 object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
