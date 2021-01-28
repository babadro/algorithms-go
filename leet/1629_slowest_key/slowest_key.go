package _1629_slowest_key

// passed. tptl best solution
func slowestKey2(releaseTimes []int, keysPressed string) byte {
	slowest, maxDuration := keysPressed[0], releaseTimes[0]

	for i := 1; i < len(keysPressed); i++ {
		duration := releaseTimes[i] - releaseTimes[i-1]
		if duration > maxDuration || (duration == maxDuration && keysPressed[i] > slowest) {
			maxDuration, slowest = duration, keysPressed[i]
		}
	}

	return slowest
}

func slowestKey(releaseTimes []int, keysPressed string) byte {
	res := [26]int{}

	for i, char := range keysPressed {
		idx := char - 'a'
		var t int
		if i > 0 {
			t = releaseTimes[i] - releaseTimes[i-1]
		} else {
			t = releaseTimes[i]
		}

		if res[idx] < t {
			res[idx] = t
		}
	}

	max, idx := 0, 0
	for i := range res {
		if res[i] >= max {
			max, idx = res[i], i
		}
	}

	return byte(idx + 'a')
}
