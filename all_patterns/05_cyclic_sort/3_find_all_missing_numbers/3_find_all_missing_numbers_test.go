package __find_all_missing_numbers

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findAllMissingNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{2, 3, 1, 8, 2, 3, 5, 1}, []int{4, 6, 7}},
		{[]int{1, 2, 1, 2, 5}, []int{3, 4}},
		{[]int{1, 1}, []int{2}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findAllMissingNumbers(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAllMissingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
