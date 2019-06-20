package stack_test

import (
	"algorithms-go/03_StacksAndQueues/02_StackArray"
	"testing"
)

func TestStack(t *testing.T) {
	stack := stack.NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if stack.Empty() {
		t.Errorf("Really stack is not empty")
	}

	popres, _ := stack.Pop()

	if popres != 3 {
		t.Errorf("Pop failed. got %v want %v", popres, 3)
	}

	stack.Pop()
	stack.Pop()

	if !stack.Empty() {
		t.Errorf("Empty is %v, want %v", stack.Empty(), true)
	}
}
