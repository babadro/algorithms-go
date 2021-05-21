package output

import (
	"io/ioutil"
	"strconv"
)

// output.txt:

// 1
// 2
// 3
// 4
// 5

func WriteIntArr(nums []int) {
	b := make([]byte, 0, len(nums))

	for i, num := range nums {
		if i > 0 {
			b = append(b, '\n')
		}

		b = strconv.AppendInt(b, int64(num), 10)
	}

	if err := ioutil.WriteFile("output.txt", b, 0644); err != nil {
		panic(err)
	}
}

// output.txt:

// abc
// def
// ghi

func WriteStrArr(arr []string) {
	b := make([]byte, 0, len(arr))

	for i, str := range arr {
		if i > 0 {
			b = append(b, '\n')
		}

		b = append(b, str...)
	}

	if err := ioutil.WriteFile("output.txt", b, 0644); err != nil {
		panic(err)
	}
}
