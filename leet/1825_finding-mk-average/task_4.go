package _1825_finding_mk_average

import "sort"

// todo 1. TLE
type MKAverage struct {
	arr        []int
	m, k       int
	calculated bool
	avgArr     []int
	avg        int
}

func Constructor(m int, k int) MKAverage {
	return MKAverage{m: m, k: k}
}

func (this *MKAverage) AddElement(num int) {
	this.arr = append(this.arr, num)
	this.calculated = false
}

func (this *MKAverage) CalculateMKAverage() int {
	if this.calculated {
		return this.avg
	}

	if len(this.arr) < this.m {
		this.avg, this.calculated = -1, true
		return this.avg
	}

	this.avgArr = make([]int, this.m)
	copy(this.avgArr, this.arr[len(this.arr)-this.m:])

	sort.Ints(this.avgArr)
	this.avgArr = this.avgArr[this.k:]
	this.avgArr = this.avgArr[:len(this.avgArr)-this.k]

	sum := 0
	for _, num := range this.avgArr {
		sum += num
	}

	res := sum / len(this.avgArr)

	this.calculated = true
	this.avg = res

	return res
}

/**
 * Your MKAverage object will be instantiated and called as such:
 * obj := Constructor(m, k);
 * obj.AddElement(num);
 * param_2 := obj.CalculateMKAverage();
 */
