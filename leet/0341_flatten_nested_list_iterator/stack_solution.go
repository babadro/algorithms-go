package _341_flatten_nested_list_iterator

import "container/list"

// todo2 understand
type NestedIterator2 struct {
	stack *list.List
}

type indexVals struct {
	vals  []*NestedInteger
	index int
}

func Constructor2(nestedList []*NestedInteger) *NestedIterator2 {
	stack := list.New()
	stack.PushBack(&indexVals{nestedList, 0})
	return &NestedIterator2{stack}
}

func (this *NestedIterator2) Next() int {
	if !this.HasNext() {
		return -1
	}
	top := this.stack.Back().Value.(*indexVals)
	val := top.vals[top.index].GetInteger()
	top.index++
	return val
}

func (this *NestedIterator2) HasNext() bool {
	stack := this.stack
	for stack.Len() > 0 {
		top := stack.Back().Value.(*indexVals)
		vals, i := top.vals, top.index
		if i >= len(vals) {
			stack.Remove(stack.Back())
		} else {
			x := vals[i]
			if x.IsInteger() {
				return true
			}
			top.index++
			stack.PushBack(&indexVals{x.GetList(), 0})
		}
	}
	this.stack = stack
	return false
}
