package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	num, arr := numAndNumArr()

	fmt.Println(num)
	fmt.Println(arr)
}

func numAndNumArr() (num int, arr []int) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	flag := false
	for scanner.Scan() {
		b := scanner.Bytes()
		curr, err := strconv.Atoi(string(b))
		if err != nil {
			panic(err)
		}

		if !flag {
			num = curr
			flag = true
		} else {
			arr = append(arr, curr)
		}

		if err = scanner.Err(); err != nil {
			panic(err)
		}
	}

	return num, arr
}
