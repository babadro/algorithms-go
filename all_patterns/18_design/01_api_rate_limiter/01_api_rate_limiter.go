package _1_api_rate_limiter

import "container/list"

func limiter(userLimit, serviceLimit, duration int, requests [][2]int) (responses []int) {
	userRequests := make(map[int]*list.List)
	serviceRequests := list.New()
	for i, req := range requests {
		userID, now := req[0], req[1]
		uRequests, ok := userRequests[userID]
		if !ok {
			uRequests = list.New()
		}

		if !check(now, userLimit, duration, uRequests) {
			responses[i] = 429
			continue
		}

		if !check(now, serviceLimit, duration, serviceRequests) {
			responses[i] = 503
			continue
		}

		uRequests.PushBack(now)
		if uRequests.Len() > userLimit {
			uRequests.Remove(uRequests.Front())
		}
		userRequests[userID] = uRequests

		serviceRequests.PushBack(now)
		if serviceRequests.Len() > serviceLimit {
			serviceRequests.Remove(serviceRequests.Front())
		}

		responses[i] = 200
	}

	return responses
}

func check(now, limit, duration int, requests *list.List) bool {
	if requests.Len() < limit || now-duration > requests.Front().Value.(int) {
		return true
	}

	return false
}
