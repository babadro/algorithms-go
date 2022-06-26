package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var userLimit, serviceLimit, duration int
	var requests = make(chan [2]int)
	var statuses = make(chan int)
	scanner := bufio.NewScanner(os.Stdin)
	flag := true
	for scanner.Scan() {
		if num, err := strconv.Atoi(scanner.Text()); err == nil && num == -1 {
			close(requests)
			close(statuses)

			break
		}

		if flag {
			if _, err := fmt.Sscanf(scanner.Text(), "%d %d %d", &userLimit, &serviceLimit, &duration); err != nil {
				panic(err)
			}

			go checkRequestsChan(userLimit, serviceLimit, duration, requests, statuses)

			flag = false

			continue
		}

		var now, userID int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &now, &userID)
		if err != nil {
			panic(err)
		}

		requests <- [2]int{now, userID}

		fmt.Println(<-statuses)
	}
}

func checkRequestsChan(userLimit, serviceLimit, duration int, reqs chan [2]int, statuses chan int) {
	var serviceRequests []int
	var usersRequests = make(map[int][]int)

	for req := range reqs {
		now, userID := req[0], req[1]

		uRequests := usersRequests[userID]
		if !check(uRequests, now, userLimit, duration) {
			statuses <- 429

			continue
		}

		if !check(serviceRequests, now, serviceLimit, duration) {
			statuses <- 503

			continue
		}

		serviceRequests = append(serviceRequests, now)
		if len(serviceRequests) > serviceLimit {
			serviceRequests = serviceRequests[1:]
		}

		uRequests = append(uRequests, now)
		if len(uRequests) > userLimit {
			uRequests = uRequests[1:]
		}

		usersRequests[userID] = uRequests

		statuses <- 200
	}
}

func check(reqs []int, now, limit, duration int) bool {
	if limit == 0 {
		return false
	}

	if len(reqs) < limit || now-duration > (reqs)[0] {
		return true
	}

	return false
}
