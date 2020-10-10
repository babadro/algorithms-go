package _341_flatten_nested_list_iterator

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		name           string
		nestedIntegers []*NestedInteger
	}{
		{
			name: "1",
			nestedIntegers: []*NestedInteger{
				{[]int{1, 1}},
				{2},
				{[]int{1, 1}},
			},
		},
		{
			name: "2",
			nestedIntegers: []*NestedInteger{
				{1},
				{2},
				{3},
			},
		},
		{
			name: "3",
			nestedIntegers: []*NestedInteger{
				{[]interface{}{[]interface{}{1, 2, []int{3, 4}}, 5}},
				{6},
				{[]int{7, 8}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iter := Constructor(tt.nestedIntegers)
			for {
				if !iter.HasNext() {
					break
				}

				t.Log(iter.Next())
			}
		})
	}
}
