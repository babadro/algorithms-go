package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	a, b := twoNums()

	out, _ := os.Create("output.txt")
	defer out.Close()

	res := a + b
	out.Write([]byte(strconv.Itoa(res)))
}

func twoNums() (a, b int) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = fmt.Fscanf(f, "%d %d", &a, &b); err != nil {
		panic(err)
	}

	return
}
