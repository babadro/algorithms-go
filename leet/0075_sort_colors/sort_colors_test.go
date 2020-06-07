package _075_sort_colors

import (
	"math/rand"
	"sort"
	"testing"
)

func TestSetColors(t *testing.T) {
	cases := [][]int{
		{1},
		{1, 2, 1, 2, 1, 2},
		{2, 1, 0, 2, 1, 0},
		{0, 1, 0, 1, 0, 1, 0, 2, 2, 0, 0, 0},
		{2, 2, 1, 1, 0, 0},
		{1, 2, 0, 2, 1, 0, 1, 1, 1, 0, 1},
	}

	seed := int64(1591505243702007200) //time.Now().UnixNano()
	source := rand.NewSource(seed)
	rnd := rand.New(source)
	for i := 0; i < 1000; i++ {
		length := rnd.Intn(40)
		arr := make([]int, length)
		for j := 0; j < length; j++ {
			arr[j] = rnd.Intn(3)
		}
		cases = append(cases, arr)
	}

	for i := range cases {
		sortedArr := make([]int, len(cases[i]))
		copy(sortedArr, cases[i])
		sortColors(sortedArr)
		if !sort.IntsAreSorted(sortedArr) {
			t.Errorf("case#%d, input %v, output %v, seed %d", i+1, cases[i], sortedArr, seed)
			break
		}
	}
}
