package _982_find_array_given_subset_sums

import (
	"reflect"
	"sort"
	"testing"
)

func Test_getSums(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, -3}, []int{-3, -2, -1, 0, 0, 1, 2, 3}},
		{[]int{1}, []int{0, 1}},
		{[]int{0, 1}, []int{0, 0, 1, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := getSums(tt.nums)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSums() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_recoverArray(t *testing.T) {
	tests := []struct {
		n    int
		sums []int
		want []int
	}{
		{1, []int{0, 1}, []int{1}},
		{3, []int{-3, -2, -1, 0, 0, 1, 2, 3}, []int{-3, 1, 2}},
		{2, []int{0, 0, 0, 0}, []int{0, 0}},
		{4, []int{0, 0, 5, 5, 4, -1, 4, 9, 9, -1, 4, 3, 4, 8, 3, 8}, []int{0, -1, 4, 5}},
		{6, []int{-755, -616, -178, -110, 199, -765, -528, 49, -1331, -1022, -377, -1249, -1321, -466, -516, -666, 10, -594, -227, -1132, 189, -538, -477, -1259, 0, -1181, -217, -487, -289, -655, 416, 127, -1032, -903, -299, -872, -239, -188, -229, -905, -367, -1142, -526, -606, -1558, -645, 39, -944, -250, -676, -456, -100, -882, -260, -843, -833, -1548, -604, -893, -954, -915, 117, 426, -1171}, []int{}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := recoverArray(tt.n, tt.sums)
			sort.Ints(got)
			sort.Ints(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("recoverArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
