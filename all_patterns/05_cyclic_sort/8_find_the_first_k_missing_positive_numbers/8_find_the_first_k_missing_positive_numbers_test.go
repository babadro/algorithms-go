package __find_the_first_k_missing_positive_numbers

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_findNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		//{[]int{3, -1, 4, 5, 5}, 3, []int{1, 2, 6}},
		{[]int{2, 3, 4}, 3, []int{1, 5, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := findNumbers(tt.nums, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
