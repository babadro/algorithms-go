package __subsets_with_duplicates

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findSubsetsWithDuplicates(t *testing.T) {
	tests := []struct {
		nums []int
		want [][]int
	}{
		{[]int{1, 3, 3}, [][]int{{}, {1}, {3}, {1, 3}, {3, 3}, {1, 3, 3}}},
		{[]int{1, 5, 3, 3}, [][]int{
			{}, {1}, {5}, {3},
			{1, 5}, {1, 3}, {5, 3}, {3, 3},
			{1, 5, 3}, {1, 3, 3}, {3, 3, 5}, {1, 5, 3, 3},
		}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			got := findSubsetsWithDuplicates(tt.nums)

			require.Equal(t, len(got), len(tt.want))

			sort.Slice(got, func(i, j int) bool {
				return less(got[i], got[j])
			})

			sort.Slice(tt.want, func(i, j int) bool {
				return less(tt.want[i], tt.want[j])
			})

			for i := range got {
				sort.Ints(got[i])
				sort.Ints(tt.want[i])
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findSubsetsWithDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func less(arr1, arr2 []int) bool {
	if len(arr1) != len(arr2) {
		return len(arr1) < len(arr2)
	}

	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return arr1[i] < arr2[i]
		}
	}

	return false
}
