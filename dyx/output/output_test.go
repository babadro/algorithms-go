package output

import (
	"fmt"
	"testing"
)

func TestWriteIntArr(t *testing.T) {
	tests := []struct {
		nums []int
	}{
		{[]int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			WriteIntArr(tt.nums)
		})
	}
}

func TestWriteStrArr(t *testing.T) {
	tests := []struct {
		arr []string
	}{
		{[]string{"abc", "efg", "hig"}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.arr), func(t *testing.T) {
			WriteStrArr(tt.arr)
		})
	}
}
