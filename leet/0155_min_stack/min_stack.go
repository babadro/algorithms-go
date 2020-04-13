package _155_min_stack

import "sort"

type MinStack struct {
	items   []item
	sortIdx int
	sorted  []int
	length  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	ind := sort.Search(this.length, func(i int) bool {
		return this.sorted[i] <= x
	})
	this.sorted = append(this.sorted, x)
	this.length++
	copy(this.sorted[ind+1:this.length], this.sorted[ind:this.length-1])
	this.sorted[ind] = x
	this.items = append(this.items, item{x, ind})
}

func (this *MinStack) Pop() {
	ind := this.items[this.length-1].sortIdx
	copy(this.sorted[ind:this.length-1], this.sorted[ind+1:this.length])
	this.sorted = this.sorted[:this.length-1]
	this.items = this.items[:this.length-1]
	this.length--
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1].value
}

func (this *MinStack) GetMin() int {
	return this.sorted[len(this.sorted)-1]
}

type item struct {
	value   int
	sortIdx int
}
