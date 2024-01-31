package _1146_snapshot_array

import (
	"testing"
)

func Test(t *testing.T) {
	a := Constructor(3)

	a.Set(0, 5)

	t.Log(a.Snap())

	a.Set(0, 6)

	res := a.Get(0, 0)

	t.Log(res)
}
