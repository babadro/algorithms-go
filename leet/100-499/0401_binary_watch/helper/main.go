package main

import (
	"fmt"
	"strings"
)

var m = make(map[int][]string)

func main() {
	for hour := 0; hour < 12; hour++ {
		hourBitCount := bitCount(hour)
		for minute := 0; minute < 60; minute++ {
			sumBitCount := hourBitCount + bitCount(minute)
			stringTime := hours[hour] + ":" + minutes[minute]
			m[sumBitCount] = append(m[sumBitCount], stringTime)
		}
	}

	var sb strings.Builder

	j := 0
	for k, v := range m {
		if j > 0 {
			sb.WriteString("},\n")
		}
		j++
		sb.WriteString(fmt.Sprintf("%d: {", k))
		for i, t := range v {
			if i > 0 {
				sb.WriteString(",")
			}
			sb.WriteString(fmt.Sprintf("%q", t))
		}
	}

	sb.WriteString("},\n")

	fmt.Println(sb.String())
}

func bitCount(num int) int {
	count := 0

	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}

	return count
}

var test = map[int][]string{
	1: {"1:01", "2:02", "3:03"},
}

var hours = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11"}

var minutes = []string{"00", "01", "02", "03", "04", "05", "06", "07", "08", "09", "10", "11", "12", "13", "14", "15", "16", "17", "18", "19", "20", "21", "22", "23", "24", "25", "26", "27", "28", "29", "30", "31", "32", "33", "34", "35", "36", "37", "38", "39", "40", "41", "42", "43", "44", "45", "46", "47", "48", "49", "50", "51", "52", "53", "54", "55", "56", "57", "58", "59"}
