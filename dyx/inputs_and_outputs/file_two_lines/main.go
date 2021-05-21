package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b string
	f, _ := os.Open("input.txt")
	fmt.Fscanf(f, "%s\n%s", &a, &b)

	fmt.Printf("%s %s", a, b)
}
