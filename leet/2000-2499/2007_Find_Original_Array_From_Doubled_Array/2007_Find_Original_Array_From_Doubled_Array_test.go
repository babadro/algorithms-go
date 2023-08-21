package _2007_Find_Original_Array_From_Doubled_Array

import (
	"reflect"
	"testing"
)

func Test_findOriginalArray(t *testing.T) {
	tests := []struct {
		changed []int
		want    []int
	}{
		{[]int{1, 3, 4, 2, 6, 8}, []int{1, 3, 4}},
		{[]int{6, 3, 0, 1}, nil},
		{[]int{1}, nil},
		{[]int{0}, nil},
		{[]int{0, 0, 0, 0}, []int{0, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := findOriginalArray(tt.changed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOriginalArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
