package __fibonacci

import "testing"

func TestFib(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(FibRec(i))
	}
}

func TestFibTopDown(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(FibTopDown(i))
	}
}

func TestFibBottomUp(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Log(FibBottomUp(i))
	}
}
