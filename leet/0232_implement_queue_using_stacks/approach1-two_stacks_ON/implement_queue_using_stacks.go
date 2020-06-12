package approach1_two_stacks_ON

type MyQueue struct {
	stack1, stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{}
}

func (this *MyQueue) Push(x int) {
	for i := len(this.stack1) - 1; i >= 0; i-- {
		this.stack2 = append(this.stack2, this.stack1[i])
	}
	this.stack1 = this.stack1[:0]
	this.stack1 = append(this.stack1, x)
	for i := len(this.stack2) - 1; i >= 0; i-- {
		this.stack1 = append(this.stack1, this.stack2[i])
	}
	this.stack2 = this.stack2[:0]
}

func (this *MyQueue) Pop() int {
	last := len(this.stack1) - 1
	res := this.stack1[last]
	this.stack1 = this.stack1[:last]
	return res
}

func (this *MyQueue) Peek() int {
	return this.stack1[len(this.stack1)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.stack1) == 0
}
