package _703_kth_largest_element_in_a_stream

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConstructor(t *testing.T) {
	tests := []struct {
		k    int
		nums []int
		args []int
		want []int
	}{
		{3, []int{4, 5, 8, 2}, []int{3, 5, 10, 9, 4}, []int{4, 5, 5, 8, 8}},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.nums), func(t *testing.T) {
			obj := Constructor(tt.k, tt.nums)
			for i := range tt.args {
				require.Equal(t, obj.Add(tt.args[i]), tt.want[i])
			}
		})
	}
}
