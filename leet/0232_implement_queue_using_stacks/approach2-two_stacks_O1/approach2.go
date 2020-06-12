package approach2_two_stacks_O1

type MyQueue struct {
	stack1, stack2 []int
	front          int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	if len(this.stack1) == 0 {
		this.front = x
	}
	this.stack1 = append(this.stack1, x)
}

func (this *MyQueue) Pop() int {
	var res int
	if len(this.stack2) > 0 {
		last := len(this.stack2) - 1
		res = this.stack2[last]
		this.stack2 = this.stack2[:last]
		return res
	}
	for i := len(this.stack1) - 1; i > 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	res = this.stack1[0]
	this.stack1 = this.stack1[:0]
	return res
}

func (this *MyQueue) Peek() int {
	if n := len(this.stack2); n > 0 {
		return this.stack2[n-1]
	}
	return this.front
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0 && len(this.stack2) == 0
}
