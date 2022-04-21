package sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMerge(t *testing.T) {
	standardCases := [][]int{
		{4, 5, 6, 1, 2, 3},
	}
	for i := range standardCases {
		mid := (len(standardCases[i]) - 1) / 2
		merge(standardCases[i], 0, mid, len(standardCases[i])-1)
		if !sort.IntsAreSorted(standardCases[i]) {
			t.Errorf("not sorted: %v", standardCases[i])
		}
		t.Log(standardCases[i])
	}
}

func TestMergeSort(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Log(seed)
	source := rand.NewSource(seed)
	rnd := rand.New(source)
	for i := 0; i < 1000; i++ {
		length := rnd.Intn(100)
		input := make([]int, length)
		for j := 0; j < length; j++ {
			input[j] = rnd.Intn(1000)
			if rnd.Intn(2) == 1 {
				input[j] *= -1
			}
		}
		mergeSort(input, 0, length-1)
		if !sort.IntsAreSorted(input) {
			t.Errorf("not sorted: %v", input)
		}
	}

	standardCases := [][]int{
		nil, {}, {0}, {1}, {-1}, {2, 1}, {1, 2},
	}
	for i := range standardCases {
		mergeSort(standardCases[i], 0, len(standardCases[i])-1)
		if !sort.IntsAreSorted(standardCases[i]) {
			t.Errorf("not sorted: %v", standardCases[i])
		}
		t.Log(standardCases[i])
	}
}

func TestMergeSortParallel(t *testing.T) {
	testFunc(t, func(nums []int) {
		mergeSortParallel(nums, 0, len(nums)-1, 0, 2)
	})
}

var input []int

func BenchmarkMergeSort(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		input = make([]int, 1_000_000)
		for j := 0; j < len(input); j++ {
			min, max := -1000, 1000
			input[j] = rand.Intn(max-min) + min
		}

		b.StartTimer()
		mergeSort(input, 0, len(input)-1)
		b.StopTimer()
	}
}

func BenchmarkMergeSortParallel(b *testing.B) {
	b.StopTimer()

	for i := 0; i < b.N; i++ {
		input = make([]int, 1_000_000)
		for j := 0; j < len(input); j++ {
			min, max := -1000, 1000
			input[j] = rand.Intn(max-min) + min
		}

		b.StartTimer()
		mergeSortParallel(input, 0, len(input)-1, 0, 2)
		b.StopTimer()
	}
}
