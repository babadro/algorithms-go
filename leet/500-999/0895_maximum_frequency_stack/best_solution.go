package _895_maximum_frequency_stack

import "container/heap"

// tptl. passed. best solution (very simple and clear)
type element struct {
	number, freq, sequenceNumber int
}

func less(el1, el2 element) bool {
	if el1.freq != el2.freq {
		return el1.freq > el2.freq
	}

	return el1.sequenceNumber > el2.sequenceNumber
}

type FreqStack2 struct {
	sequenceNum int
	h           maxHeap2
	freqMap     map[int]int
}

func Constructor2() FreqStack2 {
	return FreqStack2{
		freqMap: make(map[int]int),
	}
}

func (this *FreqStack2) Push(val int) {
	this.freqMap[val]++
	this.sequenceNum++
	heap.Push(&this.h, element{
		number:         val,
		freq:           this.freqMap[val],
		sequenceNumber: this.sequenceNum,
	})
}

func (this *FreqStack2) Pop() int {
	el := heap.Pop(&this.h).(element)

	if this.freqMap[el.number] > 1 {
		this.freqMap[el.number]--
	} else {
		delete(this.freqMap, el.number)
	}

	return el.number
}

type maxHeap2 struct {
	arr []element
}

func (h maxHeap2) Len() int { return len(h.arr) }
func (h maxHeap2) Less(i, j int) bool {
	return less(h.arr[i], h.arr[j])
}
func (h maxHeap2) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}
func (h *maxHeap2) Push(v interface{}) {
	h.arr = append(h.arr, v.(element))
}
func (h *maxHeap2) Pop() interface{} {
	last := len(h.arr) - 1
	res := h.arr[last]
	h.arr = h.arr[:last]

	return res
}
