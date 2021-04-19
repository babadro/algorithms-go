package _1825_finding_mk_average

import (
	"testing"
)

func TestConstructorBruteForce(t *testing.T) {
	obj := Constructor3(3, 1)
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

func TestMKAverage(t *testing.T) {
	tests := []struct {
		name     string
		commands []string
		input    [][]int
	}{
		{
			name:     "1",
			commands: []string{"MKAverage3", "addElement", "addElement", "calculateMKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "addElement", "calculateMKAverage", "addElement"},
			input:    [][]int{{3, 1}, {17612}, {74607}, {}, {8272}, {33433}, {}, {15456}, {64938}, {}, {99741}},
		},
	}

	var obj MKAverage3

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := range tt.commands {
				command := tt.commands[i]
				input := tt.input[i]

				switch command {
				case "MKAverage3":
					m, k := input[0], input[1]
					obj = Constructor3(m, k)
				case "addElement":
					obj.AddElement(input[0])
				case "calculateMKAverage":
					t.Log(obj.CalculateMKAverage())
				default:
					t.Errorf("unknown command %q", command)
				}
			}
		})
	}
}
