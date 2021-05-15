package tleInput

import (
	"io/ioutil"
	"testing"
)

func Test_inputConvertor(t *testing.T) {
	input := `["MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"]
[[3, 1], [3], [1], [], [10], [], [5], [5], [5], []]`

	_ = input
	b := []byte(inputConvertor(TleInputStr2))
	ioutil.WriteFile("tleInput1.go", b, 0644)
	t.Log()
}

func Test_outpotConvertor(t *testing.T) {
	output := "[null,null,null,-1,null,null,33433,null,null,33433,null]"

	t.Log(outputConvertor(output))
}
