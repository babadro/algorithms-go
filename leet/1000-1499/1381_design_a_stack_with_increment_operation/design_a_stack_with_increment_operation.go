package _1381_design_a_stack_with_increment_operation

// passed. tptl
type CustomStack []int

func Constructor(maxSize int) CustomStack {
	return make([]int, 0, maxSize)
}

func (this *CustomStack) Push(x int) {
	if len(*this) < cap(*this) {
		*this = append(*this, x)
	}
}

func (this *CustomStack) Pop() int {
	if len(*this) > 0 {
		last := len(*this) - 1
		res := (*this)[last]
		*this = (*this)[:last]
		return res
	}

	return -1
}

func (this *CustomStack) Increment(k int, val int) {
	for i := 0; i < k && i < len(*this); i++ {
		(*this)[i] += val
	}
}

/**
 * Your CustomStack object will be instantiated and called as such:
 * obj := Constructor(maxSize);
 * obj.Push(x);
 * param_2 := obj.Pop();
 * obj.Increment(k,val);
 */
