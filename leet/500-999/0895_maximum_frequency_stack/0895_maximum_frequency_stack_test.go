package _895_maximum_frequency_stack

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFreqStack(t *testing.T) {
	tests := []struct {
		ops  []interface{}
		want []interface{}
	}{
		{
			[]interface{}{nil, 5, 7, 5, 7, 4, 5, nil, nil, nil, nil},
			[]interface{}{nil, nil, nil, nil, nil, nil, nil, 5, 7, 5, 4},
		},
		//{
		//	[]interface{}{nil, 4, 0, 9, 3, 4, 2, nil, 6, nil, 1, nil, 1, nil, 4, nil, nil, nil, nil, nil, nil},
		//	[]interface{}{nil, 4, 0, 9, 3, 4, 2, nil, 6, nil, 1, nil, 1, nil, 4, nil, nil, nil, nil, nil, nil},
		//},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			require.Equal(t, len(tt.ops), len(tt.want))

			stack := Constructor2()
			for i := 1; i < len(tt.want); i++ {
				if tt.ops[i] == nil {
					// pop
					got := stack.Pop()
					want, ok := tt.want[i].(int)
					require.True(t, ok)

					_, _, ok = want, got, ok

					require.Equal(t, want, got, "want %d, got %d, idx=%d", want, got, i)
				} else {
					num, ok := tt.ops[i].(int)
					require.True(t, ok)

					stack.Push(num)
				}
			}
		})
	}
}
