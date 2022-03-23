package _2113_longest_subsring_of_one_repeating_character

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {

}

type item struct {
	idxStart  int
	idxFinish int
	char      byte
}

func (i *item) Len() int {
	return i.idxFinish - i.idxStart
}

type maxHeap struct {
	arr              []item
	numsIdxToHeapIdx map[int]int
}

func (h maxHeap) Len() int           { return len(h.arr) }
func (h maxHeap) Less(i, j int) bool { return h.arr[i].Len() > h.arr[j].Len() }

func (h maxHeap) Swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
	h.numsIdxToHeapIdx[h.arr[i].numsIdx] = i
	h.numsIdxToHeapIdx[h.arr[j].numsIdx] = j
}

func (h *maxHeap) Push(x interface{}) {
	el := x.(item)
	h.arr = append(h.arr, el)
	h.numsIdxToHeapIdx[el.numsIdx] = len(h.arr) - 1
}

func (h *maxHeap) Pop() interface{} {
	n := len(h.arr)
	res := h.arr[n-1]
	h.arr = h.arr[:n-1]

	delete(h.numsIdxToHeapIdx, res.numsIdx)

	return res
}
