package sorting

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func testFunc(t *testing.T, sortFunc func(nums []int)) {
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
			if rnd.Intn(2) == 1 {
				input[j] *= -1
			}
		}
		sortFunc(input)
		if !sort.IntsAreSorted(input) {
			t.Errorf("not sorted: %v", input)
		}
	}

	standardCases := [][]int{
		nil, {}, {0}, {1}, {-1}, {2, 1}, {1, 2},
	}
	for i := range standardCases {
		sortFunc(standardCases[i])
		if !sort.IntsAreSorted(standardCases[i]) {
			t.Errorf("not sorted: %v", standardCases[i])
		}
		t.Log(standardCases[i])
	}
}
