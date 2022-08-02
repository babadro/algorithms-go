package array

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_filterInPlace(t *testing.T) {
	tests := []struct {
		nums []int
		f    func(num int) bool
		want []int
	}{
		{[]int{1, 2, 3, 4, 1, 1, 1, 5, 6, 1}, func(num int) bool { return num != 1 }, []int{2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, func(num int) bool {
			return num%2 == 0
		}, []int{2, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			if got := filterInPlacePreserveOrder(tt.nums, tt.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("filterInPlacePreserveOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_filterInPlaceNotPreserveOrder(t *testing.T) {

	tests := []struct {
		nums []int
		f    func(num int) bool
		want []int
	}{
		{[]int{1, 2, 3, 4, 1, 1, 1, 5, 6, 1}, func(num int) bool { return num != 1 }, []int{2, 3, 4, 5, 6}},
		{[]int{1, 2, 3, 4, 5, 6}, func(num int) bool {
			return num%2 == 0
		}, []int{2, 4, 6}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := filterInPlaceNotPreserveOrder(tt.nums, tt.f)

			sort.Ints(got)
			sort.Ints(tt.want)
			require.Equal(t, tt.want, got)
		})
	}
}
