package main

import (
	"fmt"
	"os"
)

func main() {
	var first, second int64
	fmt.Fscan(os.Stdin, &first, &second)

	fmt.Println(first + second)
}
