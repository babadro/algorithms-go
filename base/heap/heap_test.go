package heap

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestMaxHeap_Build(t *testing.T) {
	tests := []struct {
		list []int
		want []int
	}{
		{
			[]int{1, 3, 5, 4, 6, 13, 10, 9, 8, 15, 17},
			[]int{17, 15, 13, 9, 6, 5, 10, 4, 8, 3, 1},
		},
		{
			[]int{},
			[]int{},
		},
		{
			[]int{1},
			[]int{1},
		},
		{
			[]int{1, 2},
			[]int{2, 1},
		},
		{
			[]int{1, 2, 3},
			[]int{3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.list), func(t *testing.T) {
			heap := MaxHeap{}
			heap.Build(tt.list)

			if !reflect.DeepEqual(heap.list, tt.want) {
				t.Errorf("got = %v, want %v", heap.list, tt.want)
			}
		})
	}
}

func TestMaxHeap_Add(t *testing.T) {
	tests := []struct {
		list []int
		want []int
	}{
		{
			[]int{2, 3, 4, 1, 5, 1, 5, 2, 10, 6},
			[]int{10, 6, 5, 5, 4, 3, 2, 2, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprint(tt.list), func(t *testing.T) {
			heap := MaxHeap{}
			for _, num := range tt.list {
				heap.Add(num)
			}

			var res []int
			for {
				if num, ok := heap.ExtractMax(); ok {
					res = append(res, num)
				} else {
					break
				}
			}

			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("got = %v, want %v", res, tt.want)
			}
		})
	}
}

func testHelper(t *testing.T, f func(t *testing.T, arr []int)) {
	seed := time.Now().UnixNano()
	t.Log(seed)
	source := rand.NewSource(seed)
	rnd := rand.New(source)
	for i := 0; i < 1000; i++ {
		length := rnd.Intn(100)
		input := make([]int, length)
		for j := 0; j < length; j++ {
			min, max := -1000, 1000
			input[j] = rnd.Intn(max-min) + min
		}

		f(t, input)
	}

	standardCases := [][]int{
		{}, {0}, {1}, {-1}, {2, 1}, {1, 2}, {1, 2, 3}, {3, 2, 1},
	}
	for i := range standardCases {
		f(t, standardCases[i])
	}
}

func TestMaxHeap_Build2(t *testing.T) {
	f := func(t *testing.T, arr []int) {
		heap := MaxHeap{}
		heap.Build(arr)
		sorted := getSortedArr(&heap)

		if !sort.SliceIsSorted(sorted, func(i, j int) bool {
			return sorted[i] > sorted[j]
		}) {
			t.Error("slice is not sorted")
		}
	}

	testHelper(t, f)
}

func TestMaxHeap_Add2(t *testing.T) {
	f := func(t *testing.T, arr []int) {
		heap := MaxHeap{}
		for _, num := range arr {
			heap.Add(num)
		}

		sorted := getSortedArr(&heap)
		sort.Slice(arr, func(i, j int) bool {
			return arr[i] > arr[j]
		})

		if !reflect.DeepEqual(sorted, arr) {
			t.Errorf("got = %v, want %v", sorted, arr)
		}
	}

	testHelper(t, f)
}

// todo 1 fails.
func TestMaxHeap_Delete(t *testing.T) {
	f := func(t *testing.T, arr []int) {
		heap := MaxHeap{}
		heap.Build(arr)

		if len(arr) > 0 {
			idx := rand.Intn(len(arr))
			heap.Delete(idx)

			sorted := getSortedArr(&heap)

			if !sort.SliceIsSorted(sorted, func(i, j int) bool {
				return sorted[i] > sorted[j]
			}) {
				t.Errorf("slice is not sorted %v", sorted)
			}

			if len(sorted) != len(arr)-1 {
				t.Error("len should be decreased by one")
			}
		}
	}

	testHelper(t, f)
}

func getSortedArr(h *MaxHeap) []int {
	sorted := make([]int, 0, h.Size())
	for {
		if num, ok := h.ExtractMax(); ok {
			sorted = append(sorted, num)
		} else {
			break
		}
	}

	return sorted
}
