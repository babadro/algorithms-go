package _911_online_election

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_binarySearch(t *testing.T) {
	tests := []struct {
		votes []int
		t     int
		want  int
	}{
		{[]int{0}, 0, 0},
		{[]int{0, 1}, 0, 0},
		{[]int{0, 1, 1, 1, 2}, 1, 3},
		{[]int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2}, 1, 9},
		{[]int{1, 2}, 0, -1},
		{[]int{5}, 4, -1},
		{[]int{1}, 2, 0},
		{[]int{1, 2}, 3, 1},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v", tt.votes), func(t *testing.T) {
			if got := binarySearch(tt.votes, tt.t); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConstructor(t *testing.T) {
	v := Constructor2([]int{0, 1, 1, 0, 0, 1, 0}, []int{0, 5, 10, 15, 20, 25, 30})

	input := []int{3, 12, 25, 15, 24, 8}
	output := []int{0, 1, 1, 0, 0, 1}

	for i := range input {
		res := v.Q(input[i])

		require.Equal(t, output[i], res)
	}
}
