package _155_min_stack

import "testing"

func TestStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(-2)
	minStack.Push(-3)
	t.Log(minStack.GetMin())
	minStack.Pop()
	t.Log(minStack.Top())
	t.Log(minStack.GetMin())
}
