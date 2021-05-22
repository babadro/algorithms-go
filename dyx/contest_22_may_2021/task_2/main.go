package main

import (
	"fmt"
	"os"
)

func main() {
	var pType string
	var left, right string
	fmt.Fscan(os.Stdin, &pType, &left, &right)

	fmt.Println(pType)
	fmt.Println(left)
	fmt.Println(right)
}
