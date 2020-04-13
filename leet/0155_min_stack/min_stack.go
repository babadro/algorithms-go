package _155_min_stack

type MinStack struct {
	items, mins []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(x int) {
	this.items = append(this.items, x)
	if len(this.mins) == 0 {
		this.mins = append(this.mins, x)
	} else {
		min := this.mins[len(this.mins)-1]
		if x <= min {
			this.mins = append(this.mins, x)
		}
	}
}

func (this *MinStack) Pop() {
	if len(this.items) == 0 {
		return
	}
	x := this.items[len(this.items)-1]
	this.items = this.items[:len(this.items)-1]
	min := this.mins[len(this.mins)-1]
	if min == x {
		this.mins = this.mins[:len(this.mins)-1]
	}
}

func (this *MinStack) Top() int {
	return this.items[len(this.items)-1]
}

func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}
