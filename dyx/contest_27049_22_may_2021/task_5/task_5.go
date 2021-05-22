package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var userLimit, serviceLimit, duration int

	_, _, _ = userLimit, serviceLimit, duration

	flag := false
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if !flag {
			input := scanner.Text()
			arr := strings.Split(input, " ")
			userLimit, _ = strconv.Atoi(arr[0])
			serviceLimit, _ = strconv.Atoi(arr[1])
			duration, _ = strconv.Atoi(arr[2])

			flag = true
		} else {
			input := scanner.Text()
			if _, err := strconv.Atoi(input); err != nil {
				break
			}

		}
	}

	fmt.Println(userLimit, " ", serviceLimit, " ", duration)
}
