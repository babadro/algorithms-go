package _470_implemend_rand10_using_rand7

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rand10(t *testing.T) {
	m := make(map[int]bool)

	for i := 0; i < 1000; i++ {
		res := rand10()

		if res < 0 {
			panic(0)
		}

		if res > 48 {
			panic(48)
		}

		m[res] = true
	}

	require.Len(t, m, 49)
}
