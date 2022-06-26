package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Условию задачи работа с вводом и выводом не соответствовала, т.е. это просто чтобы по-быстрому написать юнит-тесты
func main() {
	var userLimit, serviceLimit, duration int
	var requests [][2]int
	scanner := bufio.NewScanner(os.Stdin)
	flag := true
	for scanner.Scan() {
		if num, err := strconv.Atoi(scanner.Text()); err == nil && num == -1 {
			break
		}

		if flag {
			if _, err := fmt.Sscanf(scanner.Text(), "%d %d %d", &userLimit, &serviceLimit, &duration); err != nil {
				panic(err)
			}

			flag = false

			continue
		}

		var now, userID int
		_, err := fmt.Sscanf(scanner.Text(), "%d %d", &now, &userID)
		if err != nil {
			panic(err)
		}

		requests = append(requests, [2]int{now, userID})
	}

	statuses := checkRequests(userLimit, serviceLimit, duration, requests)

	for _, s := range statuses {
		fmt.Println(s)
	}
}

func checkRequests(userLimit, serviceLimit, duration int, reqs [][2]int) (statuses []int) {
	var serviceRequests []int
	var usersRequests = make(map[int][]int)
	statuses = make([]int, len(reqs))

	for i, req := range reqs {
		now, userID := req[0], req[1]

		uRequests := usersRequests[userID]
		if !check(uRequests, now, userLimit, duration) {
			statuses[i] = 429

			continue
		}

		if !check(serviceRequests, now, serviceLimit, duration) {
			statuses[i] = 503

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

		statuses[i] = 200
	}

	return statuses
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
