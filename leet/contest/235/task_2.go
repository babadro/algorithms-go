package _35

// todo 1
func findingUsersActiveMinutes(logs [][]int, k int) []int {
	userToUamDict := make(map[int]map[int]bool)

	for _, log := range logs {
		id, minute := log[0], log[1]

		if _, ok := userToUamDict[id]; !ok {
			userToUamDict[id] = make(map[int]bool)
		}

		userToUamDict[id][minute] = true
	}

	uamToUserCount := make(map[int]int)

	for _, uamDict := range userToUamDict {
		uam := len(uamDict)
		uamToUserCount[uam]++
	}

	res := make([]int, k)
	for i := 0; i < k; i++ {
		uam := i + 1
		usersCount := uamToUserCount[uam]

		res[i] = usersCount
	}

	return res
}
