package _118_Pascals_Triangle

import "testing"

func TestGenerate(t *testing.T) {
	cases := []int{0, 1, 2, 3, 4, 5}
	for _, c := range cases {
		t.Log(generate(c))
	}
}
