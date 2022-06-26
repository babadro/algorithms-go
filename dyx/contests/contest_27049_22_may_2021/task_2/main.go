package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const layout = "2006-01-02"
const day = 24 * time.Hour

func main() {
	inType, start, finish := parseInput()

	var res []byte
	curDay := start
	counter := 0
	for flag := true; curDay.Unix() <= finish.Unix(); curDay = curDay.Add(day) {
		if flag {
			counter++
			flag = false
			res = append(res, curDay.Format(layout)...)
			res = append(res, ' ')
		}

		if lastDayOfInterval(inType, curDay, finish) {
			flag = true
			res = append(res, curDay.Format(layout)...)
			res = append(res, '\n')
		}
	}

	counterBytes := []byte(strconv.Itoa(counter) + "\n")
	res = append(counterBytes, res...)
	if err := ioutil.WriteFile("output.txt", res, 0644); err != nil {
		panic(err)
	}
}

func lastDayOfInterval(inType string, curDay, finish time.Time) bool {
	if curDay == finish {
		return true
	}

	nextDay := curDay.Add(day)

	switch inType {
	case "WEEK":
		return nextDay.Weekday() == time.Monday
	case "MONTH":
		return nextDay.Day() == 1
	case "QUARTER":
		return nextDay.Day() == 1 &&
			(nextDay.Month() == time.January ||
				nextDay.Month() == time.April ||
				nextDay.Month() == time.July ||
				nextDay.Month() == time.October)
	case "YEAR":
		return nextDay.Day() == 1 && nextDay.Month() == time.January
	case "REVIEW":
		return nextDay.Day() == 1 && (nextDay.Month() == time.April || nextDay.Month() == time.October)
	}

	panic("Unknown interval type")
}

func parseInput() (inType string, start, finish time.Time) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var startStr, finishStr string
	if _, err = fmt.Fscanf(f, "%s\n%s %s", &inType, &startStr, &finishStr); err != nil {
		panic(err)
	}

	start, err = time.Parse(layout, startStr)
	if err != nil {
		panic(err)
	}

	finish, err = time.Parse(layout, finishStr)
	if err != nil {
		panic(err)
	}

	return
}
