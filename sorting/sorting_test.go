package sorting

import (
	"math/rand"
	"reflect"
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
		}

		makeInputCopyThenSortBothOfThemAndCompare(t, input, sortFunc)
	}

	standardCases := [][]int{
		{}, {0}, {1}, {-1}, {2, 1}, {1, 2},
	}
	for i := range standardCases {
		makeInputCopyThenSortBothOfThemAndCompare(t, standardCases[i], sortFunc)
	}
}

func makeInputCopyThenSortBothOfThemAndCompare(t *testing.T, input []int, sortFunc func(nums []int)) {
	copyArr := make([]int, len(input))
	copy(copyArr, input)
	sort.Slice(copyArr, func(i, j int) bool {
		return copyArr[i] < copyArr[j]
	})

	sortFunc(input)

	if !reflect.DeepEqual(input, copyArr) {
		t.Errorf("got = %v, want %v", input, copyArr)
	}
}
