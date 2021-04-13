package _1825_finding_mk_average

import (
	"testing"
)

func TestConstructorBruteForce(t *testing.T) {
	obj := Constructor1(3, 1)
	obj.AddElement(3)
	obj.AddElement(1)
	t.Log(obj.CalculateMKAverage())

	obj.AddElement(10)
	t.Log(obj.CalculateMKAverage())

	obj.AddElement(5)
	obj.AddElement(5)
	obj.AddElement(5)
	t.Log(obj.CalculateMKAverage())
}
