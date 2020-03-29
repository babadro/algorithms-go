package _912_sort_an_array

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestPartition(t *testing.T) {
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
		quickSort(input, 0, length-1)
		if !sort.IntsAreSorted(input) {
			t.Errorf("not sorted: %v", input)
		}
	}

	standardCases := [][]int{
		nil, {}, {0}, {1}, {-1}, {2, 1}, {1, 2},
	}
	for i := range standardCases {
		quickSort(standardCases[i], 0, len(standardCases[i])-1)
		if !sort.IntsAreSorted(standardCases[i]) {
			t.Errorf("not sorted: %v", standardCases[i])
		}
	}

}
