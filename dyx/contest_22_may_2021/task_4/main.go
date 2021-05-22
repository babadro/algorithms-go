package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
)

type Arr []int

func main() {
	//serverURL, port, a, b := parseInput()
	//
	//r := fmt.Sprintf("%s:%d?a=%d&b=%d", serverURL, port, a, b)
	//
	//resp, err := http.Get(r)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//
	//bytesResp, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	panic(err)
	//}

	//var arr []int
	//if err := json.Unmarshal(bytesResp, &arr); err != nil {
	//	panic(err)
	//}

	arr := []int{5, 4, 3, 2, 1}

	sort.Ints(arr)

	var res []byte
	for i, num := range arr {
		if i > 0 {
			res = append(res, '\n')
		}

		res = strconv.AppendInt(res, int64(num), 10)
	}

	if err := ioutil.WriteFile("output.txt", res, 0644); err != nil {
		panic(err)
	}
}

func parseInput() (serverURL string, port, a, b int) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = fmt.Fscanf(f, "%s\n%d\n%d\n%d\n", &serverURL, &port, &a, &b); err != nil {
		panic(err)
	}

	return
}
