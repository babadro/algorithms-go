package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var a, b int
	f, _ := os.Open("input.txt")
	defer f.Close()

	fmt.Fscanf(f, "%d %d", &a, &b)

	out, _ := os.Create("output.txt")
	defer out.Close()

	res := a + b
	out.Write([]byte(strconv.Itoa(res)))
}
