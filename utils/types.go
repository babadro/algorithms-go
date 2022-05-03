package utils

type IntMinHeap []int

func (h IntMinHeap) Len() int            { return len(h) }
func (h IntMinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *IntMinHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}

type IntMaxHeap []int

func (h IntMaxHeap) Len() int            { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(v interface{}) { *h = append(*h, v.(int)) }
func (h *IntMaxHeap) Pop() interface{} {
	last := len(*h) - 1
	res := (*h)[last]
	*h = (*h)[:last]

	return res
}
