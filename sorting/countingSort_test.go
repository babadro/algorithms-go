package sorting

import (
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestCounting(t *testing.T) {
	seed := time.Now().UnixNano()
	t.Log(seed)
	source := rand.NewSource(seed)
	rnd := rand.New(source)
	for i := 0; i < 1000; i++ {
		length := rnd.Intn(100)
		input := make([]byte, length)
		for j := 0; j < length; j++ {
			input[j] = byte(rnd.Intn(256))
		}

		copyArr := make([]byte, len(input))
		copy(copyArr, input)
		sort.Slice(copyArr, func(i, j int) bool {
			return copyArr[i] < copyArr[j]
		})

		Counting(input)

		if !reflect.DeepEqual(input, copyArr) {
			t.Errorf("got = %v, want %v", input, copyArr)
		}
	}

	standardCases := [][]byte{
		{}, {'a'}, {'b'}, {'a', 'b'}, {'b', 'a'},
	}
	for _, tt := range standardCases {
		t.Run(fmt.Sprintf("%s", tt), func(t *testing.T) {
			copyArr := make([]byte, len(tt))
			copy(copyArr, tt)
			sort.Slice(copyArr, func(i, j int) bool {
				return copyArr[i] < copyArr[j]
			})
			Counting(tt)
			if !reflect.DeepEqual(tt, copyArr) {
				t.Errorf("got = %s, want %s", tt, copyArr)
			}
		})
	}
}
