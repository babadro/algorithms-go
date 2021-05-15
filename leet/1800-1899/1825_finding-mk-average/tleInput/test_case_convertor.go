package tleInput

import "strings"

func inputConvertor(input string) string {
	input = strings.Replace(input, "[", "{", -1)
	input = strings.Replace(input, "]", "}", -1)

	arr := strings.Split(input, "\n")
	if len(arr) < 2 {
		panic("len must be equal 2")
	}

	commands := "[]string" + arr[0]
	arguments := "[][]int" + arr[1]

	return commands + "\n" + arguments
}

func outputConvertor(output string) string {
	output = strings.Replace(output, "[", "{", -1)
	output = strings.Replace(output, "]", "}", -1)
	output = strings.Replace(output, "null", "nil", -1)

	return "[]interface{}" + output
}
