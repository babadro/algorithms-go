package _870_minimum_speed_to_arrive_on_time

// tptl. passed. medium (hard for me)
// https://leetcode.com/problems/minimum-speed-to-arrive-on-time/discuss/1224740/Binary-Search
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if hour <= float64(n)-1 {
		return -1
	}

	l, r := 1, 10_000_000
	for l < r {
		m := (l + r) / 2
		time := 0
		for i := 0; i < n-1; i++ {
			time += dist[i] / m

			if dist[i]%m > 0 {
				time += 1
			}
		}

		fDist, fM, fTime := float64(dist[n-1]), float64(m), float64(time)

		if fDist/fM+fTime <= hour {
			r = m
		} else {
			l = m + 1
		}
	}

	return l
}
