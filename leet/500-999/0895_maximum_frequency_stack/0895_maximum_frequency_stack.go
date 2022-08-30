package _895_maximum_frequency_stack

import "container/heap"

// tptl. passed. hard.
// todo 2 investigate alternative solutions
type FreqStack struct {
	h maxHeap
}

func Constructor() FreqStack {
	return FreqStack{h: maxHeap{
		numToHeapIDx: make(map[int]int),
	}}
}

func (this *FreqStack) Push(val int) {
	heap.Push(&this.h, val)
}

func (this *FreqStack) Pop() int {
	count := len(this.h.arr[0].ids)
	count--
	this.h.arr[0].ids = this.h.arr[0].ids[:count]
	e := this.h.arr[0]
	if count == 0 {
		heap.Pop(&this.h)
	} else {
		heap.Fix(&this.h, 0)
	}

	return e.num
}

type el struct {
	num int
	ids []int
}

type maxHeap struct {
	arr          []el
	numToHeapIDx map[int]int
	idSeq        int
}

func (h maxHeap) Len() int { return len(h.arr) }
func (h maxHeap) Less(i, j int) bool {
	elI, elJ := h.arr[i], h.arr[j]
	countI, countJ := len(elI.ids), len(elJ.ids)

	if countI != countJ {
		return countI > countJ
	}

	return elI.ids[countI-1] > elJ.ids[countI-1]
}
func (h maxHeap) Swap(i, j int) {
	h.numToHeapIDx[h.arr[i].num] = j
	h.numToHeapIDx[h.arr[j].num] = i

	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *maxHeap) Push(v interface{}) {
	num := v.(int)
	idx, ok := h.numToHeapIDx[num]
	h.idSeq++
	id := h.idSeq
	if !ok {
		e := el{
			num: num,
			ids: []int{id},
		}

		heapIDx := h.Len()
		h.numToHeapIDx[num] = heapIDx
		h.arr = append(h.arr, e)

		return
	}

	h.arr[idx].ids = append(h.arr[idx].ids, id)
	heap.Fix(h, idx)
}

func (h *maxHeap) Pop() interface{} {
	last := len(h.arr) - 1
	res := h.arr[last]

	h.arr = h.arr[:last]

	delete(h.numToHeapIDx, res.num)

	return res
}
