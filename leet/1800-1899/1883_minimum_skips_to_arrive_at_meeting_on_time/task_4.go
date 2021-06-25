package _1883_minimum_skips_to_arrive_at_meeting_on_time

import "math"

// This approach is wrong. It doesn't work
func minSkipsIterative(dist []int, speed int, hoursBefore int) int {
	minTime := float64(0)
	for _, d := range dist {
		minTime += float64(d) / float64(speed)
	}

	if minTime > float64(hoursBefore) {
		return -1
	}

	n := len(dist)
	maxTime := float64(0)
	for i, d := range dist {
		if i < n-1 {
			maxTime += math.Ceil(float64(d) / float64(speed))
		} else {
			maxTime += float64(d) / float64(speed)
		}
	}

	res := 0
	for maxTime > float64(hoursBefore) {
		maxProfit, minRemaindersSum, j := float64(-1), math.MaxInt32, 0
		for i := 1; i < len(dist); i++ {
			remainder := dist[i-1] % speed

			profit := float64(0)
			if len(dist) == 2 && i == 1 {
				profit = float64(speed-remainder) / float64(speed)
			} else if remainder+dist[i]%speed <= speed && remainder > 0 && dist[i]%speed > 0 {
				profit = 1
			}

			if profit > maxProfit {
				minRemaindersSum = math.MinInt32
				maxProfit = profit
				j = i
			} else if profit == maxProfit {
				if sum := remainder + dist[i-1]%speed; sum < minRemaindersSum {
					minRemaindersSum = sum
					j = i
				}
			}
		}

		maxTime -= maxProfit

		if j > 0 {
			l, r := dist[j-1], dist[j]
			dist = append(dist[:j], dist[j+1:]...)
			dist[j-1] = l + r

			res++

			a := res
			_ = a
		}

	}

	return res
}
