package _232_implement_queue_using_stacks

// TODO 1 - implement with two stacks - read the solution and discuss exmples
type MyQueue struct {
	arr []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.arr = append(this.arr, x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	res := this.arr[0]
	this.arr = this.arr[1:]
	return res
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.arr[0]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.arr) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() int {
	res := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return res
}

func (s *stack) size() int {
	return len(*s)
}

func (s *stack) empty() bool {
	return len(*s) == 0
}
