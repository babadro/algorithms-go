package _341_flatten_nested_list_iterator

// todo 1
type NestedIterator struct {
	list []*NestedInteger
	i    int
	curr *NestedIterator
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{list: nestedList, i: 0}
}

func (this *NestedIterator) Next() int {
	res := this.list[this.i].GetInteger()
	this.i++
	return res
	/*
		if this.list[this.i].IsInteger() {
			res := this.list[this.i].GetInteger()
			this.i++
			return res
		}

		//if this.curr == nil {
		//	this.curr = Constructor(this.list[this.i].GetList())
		//}

		if !this.curr.HasNext() {
			this.i++
			this.curr = nil
			return this.Next()
		}

		return this.curr.Next()
	*/

}

// todo 1 fix infinitive loop
func (this *NestedIterator) HasNext() bool {
	for this.i < len(this.list) {
		if this.list[this.i].IsInteger() {
			return true
		}

		if this.curr == nil {
			list := this.list[this.i].GetList()
			if len(list) == 0 {
				this.i++
				continue
			}

			this.curr = Constructor(list)
		}

		if !this.curr.HasNext() {
			this.i++
			continue
		}
	}

	return false
	//if this.i == len(this.list) {
	//	return false
	//}
	//
	//if this.i == len(this.list)-1 {
	//	if this.list[this.i].IsInteger() {
	//		return true
	//	}
	//
	//	if this.curr == nil {
	//		this.curr = Constructor(this.list[this.i].GetList())
	//	}
	//
	//	return this.curr.HasNext()
	//}
	//
	//return true
}

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	val interface{}
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (this NestedInteger) IsInteger() bool {
	_, ok := this.val.(int)
	return ok
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	res, ok := this.val.(int)
	if !ok {
		return 0
	}

	return res
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (this *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	var arr []*NestedInteger
	val, ok := this.val.([]int)
	if !ok {
		interfaces, ok := this.val.([]interface{})
		if !ok {
			panic("cannot parse list as []interfaces")
		}
		for i := range val {
			arr = append(arr, &NestedInteger{val: interfaces[i]})
		}

		return arr
	}

	for i := range val {
		arr = append(arr, &NestedInteger{val: val[i]})
	}

	return arr
}
