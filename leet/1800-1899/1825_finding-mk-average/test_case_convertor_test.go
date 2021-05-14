package _1825_finding_mk_average

import "testing"

func Test_inputConvertor(t *testing.T) {
	input := `["MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"]
[[3, 1], [3], [1], [], [10], [], [5], [5], [5], []]`

	t.Log(inputConvertor(input))
}

func Test_outpotConvertor(t *testing.T) {
	output := "[null,null,null,-1,null,null,33433,null,null,33433,null]"

	t.Log(outputConvertor(output))
}
