package _0077_combinations

import (
	"testing"
)

func Test_combine(t *testing.T) {
	res := combine(5, 3)

	for _, item := range res {
		t.Log(item)
	}
}
