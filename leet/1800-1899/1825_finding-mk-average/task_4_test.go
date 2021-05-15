package _1825_finding_mk_average

import (
	"io/ioutil"
	"strconv"
	"testing"
)

func TestConstructorBruteForce(t *testing.T) {
	obj := Constructor(3, 1)
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
		want     []interface{}
	}{
		{
			name:     "1",
			commands: []string{"MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "calculateMKAverage", "addElement", "addElement", "addElement", "calculateMKAverage"},
			input:    [][]int{{3, 1}, {3}, {1}, {}, {10}, {}, {5}, {5}, {5}, {}},
			want:     []interface{}{nil, nil, nil, -1, nil, 3, nil, nil, nil, 5},
		},
		{
			name:     "2",
			commands: []string{"MKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "addElement", "calculateMKAverage", "addElement", "addElement", "calculateMKAverage", "addElement"},
			input:    [][]int{{3, 1}, {17612}, {74607}, {}, {8272}, {33433}, {}, {15456}, {64938}, {}, {99741}},
			want:     []interface{}{nil, nil, nil, -1, nil, nil, 33433, nil, nil, 33433, nil},
		},
		{
			name:     "3_tle",
			commands: tleCommands,
			input:    tleArguments,
		},
	}

	var obj MKAverage1

	for _, tt := range tests {
		var gotRes int
		output := "[]interface{}{"
		t.Run(tt.name, func(t *testing.T) {

			for i := range tt.commands {
				command := tt.commands[i]
				input := tt.input[i]
				//	want := tt.want[i]

				switch command {
				case "MKAverage":
					if len(input) != 2 {
						t.Error("wrong testcase. Len input for MKAverage should be 2")
						return
					}

					m, k := input[0], input[1]
					obj = Constructor1(m, k)
				case "addElement":
					if len(input) != 1 {
						t.Error("wrong testcase. Len input for addElement should be 1")
						return
					}

					obj.AddElement(input[0])
				case "calculateMKAverage":
					if len(input) != 0 {
						t.Error("wrong testcase. len input for calculateMKAverage should be 0")
						return
					}

					//	wantRes, ok := want.(int)
					//	if !ok {
					//		t.Error("wrong testcase. Want res should be int")
					//		return
					//	}

					gotRes = obj.CalculateMKAverage()
					//if gotRes != wantRes {
					//	t.Errorf("want %d, got %d", wantRes, gotRes)
					//}
				default:
					t.Errorf("unknown command %q", command)
					return
				}

				if i > 0 {
					output += ","
				}

				if command == "calculateMKAverage" {
					output += strconv.Itoa(gotRes)
				} else {
					output += "nil"
				}
			}
		})
		output += "}"

		if err := ioutil.WriteFile(tt.name+".txt", []byte(output), 0644); err != nil {
			panic(err)
		}
	}
}
