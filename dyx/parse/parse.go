package parse

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// input.txt:

// 2

func SingleNum() (a int) { // 2
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = fmt.Fscanf(f, "%d", &a); err != nil {
		panic(err)
	}

	return
}

// input.txt:

// 2 3

func TwoNums() (a, b int) { // 2 3
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

// input.txt:

// 5
// 1
// 0
// 1
// 0
// 1

func NumAndNumArr() (num int, arr []int) { // 5, []int{1, 0, 1, 0, 1}
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

// input.txt:

// 1
// 2
// 3
// 4
// 5

func NumArr() (arr []int) { // []int{1, 2, 3, 4, 5}
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		b := scanner.Bytes()
		curr, err := strconv.Atoi(string(b))
		if err != nil {
			panic(err)
		}

		arr = append(arr, curr)

		if err = scanner.Err(); err != nil {
			panic(err)
		}
	}

	return arr
}

// input.txt:

// abc
// def
// hig

func StringArr() (arr []string) { // []int{"abc", "def", "hig"}
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		arr = append(arr, scanner.Text())

		if err = scanner.Err(); err != nil {
			panic(err)
		}
	}

	return arr
}

// вариант, когда в стандартный ввод подается значение, и сразу ожидается вывод в консоль результата.
// И только потом подается следующее входное значение - реализован в contest_27049 task_5
